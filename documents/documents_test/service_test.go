// +build unit

package documents_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/coredocument"
	"github.com/centrifuge/go-centrifuge/anchors"
	"github.com/centrifuge/go-centrifuge/bootstrap"
	"github.com/centrifuge/go-centrifuge/bootstrap/bootstrappers/testlogging"
	"github.com/centrifuge/go-centrifuge/config"
	"github.com/centrifuge/go-centrifuge/config/configstore"
	"github.com/centrifuge/go-centrifuge/contextutil"
	"github.com/centrifuge/go-centrifuge/documents"
	"github.com/centrifuge/go-centrifuge/documents/invoice"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/ethereum"
	"github.com/centrifuge/go-centrifuge/identity"
	"github.com/centrifuge/go-centrifuge/protobufs/gen/go/document"
	"github.com/centrifuge/go-centrifuge/storage/leveldb"
	"github.com/centrifuge/go-centrifuge/testingutils/commons"
	"github.com/centrifuge/go-centrifuge/testingutils/config"
	"github.com/centrifuge/go-centrifuge/testingutils/documents"
	"github.com/centrifuge/go-centrifuge/testingutils/identity"
	"github.com/centrifuge/go-centrifuge/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var testRepoGlobal documents.Repository
var (
	did       = testingidentity.GenerateRandomDID()
	accountID = did[:]
)

var ctx = map[string]interface{}{}
var cfg config.Configuration

func TestMain(m *testing.M) {
	ethClient := &testingcommons.MockEthClient{}
	ethClient.On("GetEthClient").Return(nil)
	ctx[ethereum.BootstrappedEthereumClient] = ethClient
	ibootstrappers := []bootstrap.TestBootstrapper{
		&testlogging.TestLoggingBootstrapper{},
		&config.Bootstrapper{},
	}
	bootstrap.RunTestBootstrappers(ibootstrappers, ctx)
	cfg = ctx[bootstrap.BootstrappedConfig].(config.Configuration)
	cfg.Set("identityId", did.String())
	result := m.Run()
	bootstrap.RunTestTeardown(ibootstrappers)
	os.Exit(result)
}

func TestService_ReceiveAnchoredDocument(t *testing.T) {
	srv := documents.DefaultService(cfg, nil, nil, documents.NewServiceRegistry(), nil)

	// self failed
	err := srv.ReceiveAnchoredDocument(context.Background(), nil, did)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentConfigAccountID, err))

	// nil model
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	acc, err := contextutil.Account(ctxh)
	assert.NoError(t, err)
	err = srv.ReceiveAnchoredDocument(ctxh, nil, did)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))

	// first version of the document but not saved
	id2 := testingidentity.GenerateRandomDID()
	doc, _ := createCDWithEmbeddedInvoice(t, ctxh, []identity.DID{id2}, true)
	idSrv := new(testingcommons.MockIdentityService)
	idSrv.On("ValidateSignature", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ar := new(mockAnchorRepo)
	docRoot, err := doc.CalculateDocumentRoot()
	assert.NoError(t, err)
	dr, err := anchors.ToDocumentRoot(docRoot)
	assert.NoError(t, err)
	zeros := [32]byte{}
	zeroRoot, err := anchors.ToDocumentRoot(zeros[:])
	assert.NoError(t, err)
	nextAid, err := anchors.ToAnchorID(doc.NextVersion())
	ar.On("GetAnchorData", nextAid).Return(zeroRoot, time.Now(), errors.New("missing"))
	ar.On("GetAnchorData", mock.Anything).Return(dr, time.Now(), nil)
	srv = documents.DefaultService(cfg, testRepo(), ar, documents.NewServiceRegistry(), idSrv)
	err = srv.ReceiveAnchoredDocument(ctxh, doc, did)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentPersistence, err))
	ar.AssertExpectations(t)
	idSrv.AssertExpectations(t)

	// new document with saved
	doc, _ = createCDWithEmbeddedInvoice(t, ctxh, []identity.DID{id2}, false)
	idSrv = new(testingcommons.MockIdentityService)
	idSrv.On("ValidateSignature", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ar = new(mockAnchorRepo)
	docRoot, err = doc.CalculateDocumentRoot()
	assert.NoError(t, err)
	dr, err = anchors.ToDocumentRoot(docRoot)
	assert.NoError(t, err)
	nextAid, err = anchors.ToAnchorID(doc.NextVersion())
	assert.NoError(t, err)
	ar.On("GetAnchorData", nextAid).Return(zeroRoot, time.Now(), errors.New("missing"))
	ar.On("GetAnchorData", mock.Anything).Return(dr, time.Now(), nil)
	srv = documents.DefaultService(cfg, testRepo(), ar, documents.NewServiceRegistry(), idSrv)
	err = srv.ReceiveAnchoredDocument(ctxh, doc, did)
	assert.NoError(t, err)
	ar.AssertExpectations(t)
	idSrv.AssertExpectations(t)

	// prepare a new version
	err = doc.AddNFT(true, testingidentity.GenerateRandomDID().ToAddress(), utils.RandomSlice(32))
	assert.NoError(t, err)
	err = doc.AddUpdateLog(did)
	assert.NoError(t, err)
	_, err = doc.CalculateDataRoot()
	assert.NoError(t, err)
	sr, err := doc.CalculateSigningRoot()
	assert.NoError(t, err)
	sig, err := acc.SignMsg(sr)
	assert.NoError(t, err)

	doc.AppendSignatures(sig)
	ndr, err := doc.CalculateDocumentRoot()
	assert.NoError(t, err)
	err = testRepo().Create(did[:], doc.CurrentVersion(), doc)
	assert.NoError(t, err)

	// invalid transition for id3
	id3 := testingidentity.GenerateRandomDID()
	err = srv.ReceiveAnchoredDocument(ctxh, doc, id3)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalid, err))
	assert.Contains(t, err.Error(), "invalid document state transition")

	// valid transition for id2
	idSrv.On("ValidateSignature", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	ar = new(mockAnchorRepo)
	dr, err = anchors.ToDocumentRoot(ndr)
	assert.NoError(t, err)
	nextAid, err = anchors.ToAnchorID(doc.NextVersion())
	assert.NoError(t, err)
	ar.On("GetAnchorData", nextAid).Return(zeroRoot, time.Now(), errors.New("missing"))
	ar.On("GetAnchorData", mock.Anything).Return(dr, time.Now(), nil)

	srv = documents.DefaultService(cfg, testRepo(), ar, documents.NewServiceRegistry(), idSrv)
	err = srv.ReceiveAnchoredDocument(ctxh, doc, id2)
	assert.NoError(t, err)
	ar.AssertExpectations(t)
	idSrv.AssertExpectations(t)
}

func getServiceWithMockedLayers() (documents.Service, testingcommons.MockIdentityService) {
	repo := testRepo()
	idService := testingcommons.MockIdentityService{}
	idService.On("ValidateSignature", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	mockAnchor = &mockAnchorRepo{}
	return documents.DefaultService(cfg, repo, mockAnchor, documents.NewServiceRegistry(), &idService), idService
}

type mockAnchorRepo struct {
	mock.Mock
	anchors.AnchorRepository
}

var mockAnchor *mockAnchorRepo

func (r *mockAnchorRepo) GetAnchorData(anchorID anchors.AnchorID) (docRoot anchors.DocumentRoot, anchoredTime time.Time, err error) {
	args := r.Called(anchorID)
	docRoot, _ = args.Get(0).(anchors.DocumentRoot)
	anchoredTime, _ = args.Get(1).(time.Time)
	return docRoot, anchoredTime, args.Error(2)
}

// Functions returns service mocks
func mockSignatureCheck(t *testing.T, i *invoice.Invoice, idService testingcommons.MockIdentityService) testingcommons.MockIdentityService {
	anchorID, _ := anchors.ToAnchorID(i.ID())
	dr, err := i.CalculateDocumentRoot()
	assert.NoError(t, err)
	docRoot, err := anchors.ToDocumentRoot(dr)
	assert.NoError(t, err)
	mockAnchor.On("GetAnchorData", anchorID).Return(docRoot, time.Now(), nil)
	nextAid, err := anchors.ToAnchorID(i.NextVersion())
	assert.NoError(t, err)
	zeros := [32]byte{}
	zeroRoot, err := anchors.ToDocumentRoot(zeros[:])
	mockAnchor.On("GetAnchorData", nextAid).Return(zeroRoot, time.Now(), errors.New("missing"))
	return idService
}

func TestService_CreateProofs(t *testing.T) {
	service, idService := getServiceWithMockedLayers()
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	i, _ := createCDWithEmbeddedInvoice(t, ctxh, nil, false)
	idService = mockSignatureCheck(t, i.(*invoice.Invoice), idService)
	proof, err := service.CreateProofs(ctxh, i.ID(), []string{"invoice.number"})
	assert.Nil(t, err)
	assert.Equal(t, i.ID(), proof.DocumentID)
	assert.Equal(t, i.CurrentVersion(), proof.VersionID)
	assert.Equal(t, len(proof.FieldProofs), 1)
	assert.Equal(t, proof.FieldProofs[0].GetCompactName(), []byte{0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1})
}

func TestService_CreateProofsInvalidField(t *testing.T) {
	service, idService := getServiceWithMockedLayers()
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	i, _ := createCDWithEmbeddedInvoice(t, ctxh, nil, false)
	idService = mockSignatureCheck(t, i.(*invoice.Invoice), idService)
	_, err := service.CreateProofs(ctxh, i.CurrentVersion(), []string{"invalid_field"})
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentProof, err))
}

func TestService_CreateProofsDocumentDoesntExist(t *testing.T) {
	service, _ := getServiceWithMockedLayers()
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	_, err := service.CreateProofs(ctxh, utils.RandomSlice(32), []string{"invoice.number"})
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNotFound, err))
}

func TestService_CreateProofsForVersion(t *testing.T) {
	service, idService := getServiceWithMockedLayers()
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	i, _ := createCDWithEmbeddedInvoice(t, ctxh, nil, false)
	idService = mockSignatureCheck(t, i.(*invoice.Invoice), idService)
	proof, err := service.CreateProofsForVersion(ctxh, i.ID(), i.CurrentVersion(), []string{"invoice.number"})
	assert.Nil(t, err)
	assert.Equal(t, i.ID(), proof.DocumentID)
	assert.Equal(t, i.CurrentVersion(), proof.VersionID)
	assert.Equal(t, len(proof.FieldProofs), 1)
	assert.Equal(t, proof.FieldProofs[0].GetCompactName(), []byte{0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1})
}

func TestService_RequestDocumentSignature(t *testing.T) {
	srv, _ := getServiceWithMockedLayers()

	mockAnchor.On("GetAnchorData", mock.Anything).Return(nil, nil, errors.New("missing"))
	// self failed
	_, err := srv.RequestDocumentSignature(context.Background(), nil, did)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentConfigAccountID, err))

	// nil model
	tc, err := configstore.NewAccount("main", cfg)
	assert.NoError(t, err)
	acc := tc.(*configstore.Account)
	acc.IdentityID = did[:]
	ctxh, err := contextutil.New(context.Background(), acc)
	assert.NoError(t, err)
	_, err = srv.RequestDocumentSignature(ctxh, nil, did)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentNil, err))

	// add doc to repo
	id := testingidentity.GenerateRandomDID()
	doc, _ := createCDWithEmbeddedInvoice(t, ctxh, []identity.DID{id}, false)
	idSrv := new(testingcommons.MockIdentityService)
	idSrv.On("ValidateSignature", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	srv = documents.DefaultService(cfg, testRepo(), mockAnchor, documents.NewServiceRegistry(), idSrv)

	// prepare a new version
	err = doc.AddNFT(true, testingidentity.GenerateRandomDID().ToAddress(), utils.RandomSlice(32))
	assert.NoError(t, err)
	err = doc.AddUpdateLog(did)
	_, err = doc.CalculateDataRoot()
	assert.NoError(t, err)
	sr, err := doc.CalculateSigningRoot()
	assert.NoError(t, err)
	sig, err := acc.SignMsg(sr)
	assert.NoError(t, err)

	doc.AppendSignatures(sig)
	_, err = doc.CalculateDocumentRoot()
	assert.NoError(t, err)

	// invalid transition
	id2 := testingidentity.GenerateRandomDID()
	_, err = srv.RequestDocumentSignature(ctxh, doc, id2)
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentInvalid, err))
	assert.Contains(t, err.Error(), "invalid document state transition")

	// valid transition
	_, err = srv.RequestDocumentSignature(ctxh, doc, did)
	assert.NoError(t, err)
}

func TestService_CreateProofsForVersionDocumentDoesntExist(t *testing.T) {
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	i, _ := createCDWithEmbeddedInvoice(t, ctxh, nil, false)
	s, _ := getServiceWithMockedLayers()
	_, err := s.CreateProofsForVersion(ctxh, i.ID(), utils.RandomSlice(32), []string{"invoice.number"})
	assert.Error(t, err)
	assert.True(t, errors.IsOfType(documents.ErrDocumentVersionNotFound, err))
}

func TestService_GetCurrentVersion_successful(t *testing.T) {
	service, _ := getServiceWithMockedLayers()
	documentIdentifier := utils.RandomSlice(32)
	const amountVersions = 10

	version := documentIdentifier
	var currentVersion []byte

	nonExistingVersion := utils.RandomSlice(32)
	for i := 0; i < amountVersions; i++ {

		var next []byte
		if i != amountVersions-1 {
			next = utils.RandomSlice(32)
		} else {
			next = nonExistingVersion
		}

		cd := coredocumentpb.CoreDocument{
			DocumentIdentifier: documentIdentifier,
			CurrentVersion:     version,
			NextVersion:        next,
		}

		ga := new(documents.Decimal)
		assert.NoError(t, ga.SetString(fmt.Sprint(i+1)))
		inv := &invoice.Invoice{
			GrossAmount:  ga,
			CoreDocument: documents.NewCoreDocumentFromProtobuf(cd),
		}

		err := testRepo().Create(accountID, version, inv)
		currentVersion = version
		version = next
		assert.Nil(t, err)
	}

	ctxh := testingconfig.CreateAccountContext(t, cfg)
	model, err := service.GetCurrentVersion(ctxh, documentIdentifier)
	assert.Nil(t, err)
	assert.Equal(t, model.CurrentVersion(), currentVersion, "should return latest version")
	assert.Equal(t, model.NextVersion(), nonExistingVersion, "latest version should have a non existing id as nextVersion ")

}

func TestService_GetVersion_successful(t *testing.T) {
	service, _ := getServiceWithMockedLayers()
	documentIdentifier := utils.RandomSlice(32)
	currentVersion := utils.RandomSlice(32)
	cd := coredocumentpb.CoreDocument{
		DocumentIdentifier: documentIdentifier,
		CurrentVersion:     currentVersion,
	}
	ga := new(documents.Decimal)
	assert.NoError(t, ga.SetString("60"))
	inv := &invoice.Invoice{
		GrossAmount:  ga,
		CoreDocument: documents.NewCoreDocumentFromProtobuf(cd),
	}

	ctxh := testingconfig.CreateAccountContext(t, cfg)
	err := testRepo().Create(accountID, currentVersion, inv)
	assert.Nil(t, err)

	mod, err := service.GetVersion(ctxh, documentIdentifier, currentVersion)
	assert.Nil(t, err)

	assert.Equal(t, documentIdentifier, mod.ID(), "should be same document Identifier")
	assert.Equal(t, currentVersion, mod.CurrentVersion(), "should be same version")
}

func TestService_GetCurrentVersion_error(t *testing.T) {
	service, _ := getServiceWithMockedLayers()
	documentIdentifier := utils.RandomSlice(32)

	//document is not existing
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	_, err := service.GetCurrentVersion(ctxh, documentIdentifier)
	assert.True(t, errors.IsOfType(documents.ErrDocumentVersionNotFound, err))

	cd := coredocumentpb.CoreDocument{
		DocumentIdentifier: documentIdentifier,
		CurrentVersion:     documentIdentifier,
	}

	ga := new(documents.Decimal)
	assert.NoError(t, ga.SetString("60"))
	inv := &invoice.Invoice{
		GrossAmount:  ga,
		CoreDocument: documents.NewCoreDocumentFromProtobuf(cd),
	}

	err = testRepo().Create(accountID, documentIdentifier, inv)
	assert.Nil(t, err)

	_, err = service.GetCurrentVersion(ctxh, documentIdentifier)
	assert.Nil(t, err)
}

func TestService_GetVersion_error(t *testing.T) {
	service, _ := getServiceWithMockedLayers()

	documentIdentifier := utils.RandomSlice(32)
	currentVersion := utils.RandomSlice(32)

	//document is not existing
	ctxh := testingconfig.CreateAccountContext(t, cfg)
	_, err := service.GetVersion(ctxh, documentIdentifier, currentVersion)
	assert.True(t, errors.IsOfType(documents.ErrDocumentVersionNotFound, err))

	cd := coredocumentpb.CoreDocument{
		DocumentIdentifier: documentIdentifier,
		CurrentVersion:     currentVersion,
	}
	ga := new(documents.Decimal)
	assert.NoError(t, ga.SetString("60"))
	inv := &invoice.Invoice{
		GrossAmount:  ga,
		CoreDocument: documents.NewCoreDocumentFromProtobuf(cd),
	}
	err = testRepo().Create(accountID, currentVersion, inv)
	assert.Nil(t, err)

	//random version
	_, err = service.GetVersion(ctxh, documentIdentifier, utils.RandomSlice(32))
	assert.True(t, errors.IsOfType(documents.ErrDocumentVersionNotFound, err))

	//random document id
	_, err = service.GetVersion(ctxh, utils.RandomSlice(32), documentIdentifier)
	assert.True(t, errors.IsOfType(documents.ErrDocumentVersionNotFound, err))
}

func testRepo() documents.Repository {
	if testRepoGlobal != nil {
		return testRepoGlobal
	}

	ldb, err := leveldb.NewLevelDBStorage(leveldb.GetRandomTestStoragePath())
	if err != nil {
		panic(err)
	}
	testRepoGlobal = documents.NewDBRepository(leveldb.NewLevelDBRepository(ldb))
	testRepoGlobal.Register(&invoice.Invoice{})
	return testRepoGlobal
}

func TestService_Exists(t *testing.T) {
	service, _ := getServiceWithMockedLayers()
	documentIdentifier := utils.RandomSlice(32)
	ctxh := testingconfig.CreateAccountContext(t, cfg)

	//document is not existing
	_, err := service.GetCurrentVersion(ctxh, documentIdentifier)
	assert.True(t, errors.IsOfType(documents.ErrDocumentVersionNotFound, err))

	cd := coredocumentpb.CoreDocument{
		DocumentIdentifier: documentIdentifier,
		CurrentVersion:     documentIdentifier,
	}
	ga := new(documents.Decimal)
	assert.NoError(t, ga.SetString("60"))
	inv := &invoice.Invoice{
		GrossAmount:  ga,
		CoreDocument: documents.NewCoreDocumentFromProtobuf(cd),
	}

	err = testRepo().Create(accountID, documentIdentifier, inv)
	assert.NoError(t, err)

	exists := service.Exists(ctxh, documentIdentifier)
	assert.True(t, exists, "document should exist")

	exists = service.Exists(ctxh, utils.RandomSlice(32))
	assert.False(t, exists, "document should not exist")

}

func createCDWithEmbeddedInvoice(t *testing.T, ctx context.Context, collaborators []identity.DID, skipSave bool) (documents.Model, coredocumentpb.CoreDocument) {
	i := new(invoice.Invoice)
	data := testingdocuments.CreateInvoicePayload()
	if len(collaborators) > 0 {
		var cs []string
		for _, c := range collaborators {
			cs = append(cs, c.String())
		}
		data.WriteAccess = &documentpb.WriteAccess{Collaborators: cs}
	}

	err := i.InitInvoiceInput(data, did)
	assert.NoError(t, err)
	err = i.AddUpdateLog(did)
	assert.NoError(t, err)
	_, err = i.CalculateDataRoot()
	assert.NoError(t, err)
	sr, err := i.CalculateSigningRoot()
	assert.NoError(t, err)

	acc, err := contextutil.Account(ctx)
	assert.NoError(t, err)

	sig, err := acc.SignMsg(sr)
	assert.NoError(t, err)

	i.AppendSignatures(sig)
	_, err = i.CalculateDocumentRoot()
	assert.NoError(t, err)
	cd, err := i.PackCoreDocument()
	assert.NoError(t, err)

	if !skipSave {
		err = testRepo().Create(accountID, i.CurrentVersion(), i)
		assert.NoError(t, err)
	}
	return i, cd
}
