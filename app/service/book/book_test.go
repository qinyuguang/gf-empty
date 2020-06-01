package book_test

import (
	"context"
	"errors"

	dto_book "gf-empty/app/dto/book"
	mock_dto_book "gf-empty/app/dto/book/mock"
	model_book "gf-empty/app/model/book"
	mock_model_book "gf-empty/app/model/book/mock"
	service_book "gf-empty/app/service/book"
	custom_errors "gf-empty/library/errors"
	bookpb "gf-empty/proto/message/book"

	"github.com/gogf/gf/os/gtime"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	var (
		ctx           context.Context
		ctrl          *gomock.Controller
		mockModelBook *mock_model_book.MockIModel
		mockDTOBook   *mock_dto_book.MockRequest
		entity        *model_book.Entity
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockModelBook = mock_model_book.NewMockIModel(ctrl)
		model_book.Ins = mockModelBook

		entity = &model_book.Entity{
			Id:       1,
			Name:     "book1",
			Author:   "author1",
			Price:    12.00,
			CreateAt: gtime.Now(),
			UpdateAt: gtime.Now(),
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Create", func() {
		var (
			requestRPC dto_book.Request
		)

		BeforeEach(func() {
			requestRPC = &dto_book.RPC{
				Create: &bookpb.CreateRequest{
					Book: &bookpb.Book{
						Name:   "book1",
						Author: "author1",
						Price:  12.00,
					},
				},
			}
		})

		Context("without error", func() {
			It("should create success from rpc", func() {
				mockModelBook.EXPECT().Create(gomock.Any()).Return(nil)

				err := service_book.Ins.Create(ctx, requestRPC)

				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("with error", func() {
			It("when wrong dto", func() {
				mockDTOBook = mock_dto_book.NewMockRequest(ctrl)
				mockDTOBook.EXPECT().CreateRequest().Return(nil, custom_errors.NewInternalError(custom_errors.UNAVAILABLE, "error"))

				err := service_book.Ins.Create(ctx, mockDTOBook)

				Expect(err).To(HaveOccurred())
				Expect(err.Code()).To(Equal(custom_errors.UNAVAILABLE))
			})

			It("when create", func() {
				mockModelBook.EXPECT().Create(gomock.Any()).Return(errors.New("error"))

				err := service_book.Ins.Create(ctx, requestRPC)

				Expect(err).To(HaveOccurred())
				Expect(err.Code()).To(Equal(custom_errors.UNAVAILABLE))
			})
		})
	})

	Describe("Get", func() {
		var (
			requestRPC dto_book.Request
		)

		BeforeEach(func() {
			requestRPC = &dto_book.RPC{
				Get: &bookpb.GetRequest{
					Id: 1,
				},
			}
		})

		Context("without error", func() {
			It("should exist", func() {
				mockModelBook.EXPECT().GetByID(gomock.Any()).Return(entity, nil)

				res, err := service_book.Ins.Get(ctx, requestRPC)

				Expect(res).To(Equal(entity))
				Expect(err).NotTo(HaveOccurred())
			})

			It("should not exist", func() {
				mockModelBook.EXPECT().GetByID(gomock.Any()).Return(nil, nil)

				res, err := service_book.Ins.Get(ctx, requestRPC)

				Expect(res).To(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("with error", func() {
			It("when wrong dto", func() {
				mockDTOBook = mock_dto_book.NewMockRequest(ctrl)
				mockDTOBook.EXPECT().GetRequest().Return(nil, custom_errors.NewInternalError(custom_errors.UNAVAILABLE, "error"))

				res, err := service_book.Ins.Get(ctx, mockDTOBook)

				Expect(res).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Code()).To(Equal(custom_errors.UNAVAILABLE))
			})

			It("when create", func() {
				mockModelBook.EXPECT().GetByID(gomock.Any()).Return(nil, errors.New("error"))

				res, err := service_book.Ins.Get(ctx, requestRPC)

				Expect(res).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err.Code()).To(Equal(custom_errors.UNAVAILABLE))
			})
		})
	})
})
