## ShopOps API Documentation

## https://documenter.getpostman.com/view/48995984/2sBXc7JizX

This document provides a concise reference of the HTTP API implemented in `Delivery/routers/router.go`.

Base path: `/api/v1`

Authentication: protected endpoints require a valid JWT in the `Authorization: Bearer <token>` header.

Public endpoints

- POST /api/v1/auth/register  Register a new user
- POST /api/v1/auth/login  Authenticate and obtain JWT
- POST /api/v1/auth/refresh  Refresh access token

Protected endpoints (require authentication)

- Users
  - GET /api/v1/users/me  Get current authenticated user
  - PATCH /api/v1/users/me  Update current authenticated user

- Businesses
  - POST /api/v1/businesses  Create a new business
  - GET /api/v1/businesses  List businesses
  - GET /api/v1/businesses/:businessId  Get business details
  - PATCH /api/v1/businesses/:businessId  Update business

Business-scoped routes (require `:businessId` and Business middleware)
All routes below are mounted under `/api/v1/businesses/:businessId` and require business membership.

- Sales
  - POST /api/v1/businesses/:businessId/sales  Create a sale
  - GET /api/v1/businesses/:businessId/sales  List sales
  - GET /api/v1/businesses/:businessId/sales/summary  Sales summary
  - GET /api/v1/businesses/:businessId/sales/stats  Sales statistics
  - GET /api/v1/businesses/:businessId/sales/:saleId  Get sale
  - PATCH /api/v1/businesses/:businessId/sales/:saleId  Update sale
  - DELETE /api/v1/businesses/:businessId/sales/:saleId  Void/delete sale

- Expenses
  - POST /api/v1/businesses/:businessId/expenses  Create an expense
  - GET /api/v1/businesses/:businessId/expenses  List expenses
  - GET /api/v1/businesses/:businessId/expenses/summary  Expense summary
  - GET /api/v1/businesses/:businessId/expenses/categories  List expense categories
  - GET /api/v1/businesses/:businessId/expenses/:expenseId  Get expense
  - PATCH /api/v1/businesses/:businessId/expenses/:expenseId  Update expense
  - DELETE /api/v1/businesses/:businessId/expenses/:expenseId  Void/delete expense

- Inventory (products)
  - POST /api/v1/businesses/:businessId/inventory/products  Create a product
  - GET /api/v1/businesses/:businessId/inventory/products  List products
  - GET /api/v1/businesses/:businessId/inventory/products/low-stock  List low-stock products
  - GET /api/v1/businesses/:businessId/inventory/products/:productId  Get product
  - PATCH /api/v1/businesses/:businessId/inventory/products/:productId  Update product
  - DELETE /api/v1/businesses/:businessId/inventory/products/:productId  Delete product
  - POST /api/v1/businesses/:businessId/inventory/products/:productId/adjust  Adjust product stock
  - GET /api/v1/businesses/:businessId/inventory/products/:productId/history  Product stock history

- Reports
  - GET /api/v1/businesses/:businessId/reports/dashboard  Dashboard report
  - GET /api/v1/businesses/:businessId/reports/sales  Sales report
  - GET /api/v1/businesses/:businessId/reports/expenses  Expenses report
  - GET /api/v1/businesses/:businessId/reports/profit  Profit report
  - GET /api/v1/businesses/:businessId/reports/inventory  Inventory report
  - GET /api/v1/businesses/:businessId/reports/export  Export report (download/export)
  - GET /api/v1/businesses/:businessId/reports/profit/summary  Profit summary
  - GET /api/v1/businesses/:businessId/reports/profit/trends  Profit trends

- Sync
  - POST /api/v1/businesses/:businessId/sync/batch  Process a sync batch
  - GET /api/v1/businesses/:businessId/sync/status  Get sync status

Notes

- The router applies a CORS middleware allowing common headers and methods.
- Business-scoped routes use `Infrastructure.BusinessMiddleware()` to validate business access.
- All protected routes are grouped under `/api/v1` and use the configured `AuthMiddleware`.
- For details about request/response bodies, consult the individual controller implementations in `Delivery/controllers`.

If you want, I can extend this doc with example request/response schemas for the most-used endpoints.
