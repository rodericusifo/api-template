// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"category-service/internal/domain/core/category/service"

	pb_category "category-service/internal/proto/category"
	internal_registry_core_category_service "category-service/internal/registry/core/category/service"
)

type CategoryHandler struct {
	pb_category.UnimplementedCategoryHandlerServer
	CategoryService service.ICategoryService
}

func InitCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		UnimplementedCategoryHandlerServer: pb_category.UnimplementedCategoryHandlerServer{},
		CategoryService:                    internal_registry_core_category_service.CategoryService(),
	}
}
