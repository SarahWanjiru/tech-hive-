// API Response Types
export interface GeneralResponse<T = any> {
  code: number
  message: string
  data: T
}

// User Types
export interface User {
  id: number
  name: string
  email: string
  role: 'admin' | 'customer'
  created_at: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface RegisterRequest {
  name: string
  email: string
  password: string
  role: 'admin' | 'customer'
}

export interface AuthResponse {
  token: string
  user: User
}

// Product Types
export interface Product {
  id: number
  product_id: string
  name: string
  description: string
  price: number
  stock: number
  image_url?: string
  created_at: string
}

export interface ProductCreateOrUpdateModel {
  name: string
  description: string
  price: number
  stock: number
  image_url?: string
}

export interface ProductSearchModel {
  name?: string
  min_price?: number
  max_price?: number
  in_stock?: boolean
  page: number
  limit: number
  sort_by?: string
  sort_order?: 'asc' | 'desc'
}

export interface ProductSearchResponse {
  products: Product[]
  total_count: number
  page: number
  limit: number
}

// Cart Types
export interface CartItem {
  id: number
  cart_id: number
  product_id: string
  product: Product
  quantity: number
  price: number
  created_at: string
}

export interface Cart {
  id: number
  user_id: number
  items: CartItem[]
  total: number
  created_at: string
}

export interface AddToCartModel {
  product_id: string
  quantity: number
}

export interface UpdateCartItemModel {
  quantity: number
}

// Order Types
export interface OrderItem {
  id: number
  order_id: number
  product_id: string
  product: Product
  quantity: number
  price: number
  created_at: string
}

export interface Payment {
  id: number
  order_id: number
  transaction_id: string
  status: 'pending' | 'success' | 'failed'
  paid_at?: string
}

export interface Order {
  id: number
  user_id: number
  total: number
  status: 'pending' | 'confirmed' | 'shipped' | 'delivered' | 'cancelled'
  created_at: string
  order_items: OrderItem[]
  payment?: Payment
}

export interface CreateOrderModel {
  cart_id: number
  shipping_address: string
}

export interface UpdateOrderStatusModel {
  status: string
}

// M-Pesa Types
export interface MpesaPaymentRequest {
  order_id: number
  phone_number: string
  amount: number
}

export interface MpesaSTKPushRequest {
  business_short_code: string
  password: string
  timestamp: string
  transaction_type: string
  amount: number
  party_a: string
  party_b: string
  phone_number: string
  call_back_url: string
  account_reference: string
  transaction_desc: string
}

export interface MpesaPaymentResponse {
  merchant_request_id: string
  checkout_request_id: string
  response_code: string
  response_message: string
  customer_message: string
}

export interface MpesaCallbackMetadataItem {
  name: string
  value: string
}

export interface MpesaCallbackMetadata {
  item: MpesaCallbackMetadataItem[]
}

export interface MpesaCallbackRequest {
  merchant_request_id: string
  checkout_request_id: string
  result_code: number
  result_desc: string
  callback_metadata?: MpesaCallbackMetadata
}

// API Error Types
export interface ApiError {
  code: number
  message: string
  details?: string
}

// Pagination Types
export interface PaginationParams {
  page: number
  limit: number
  sort_by?: string
  sort_order?: 'asc' | 'desc'
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  limit: number
  total_pages: number
}