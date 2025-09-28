import { apiClient } from '../api/client'
import type {
  MpesaPaymentRequest,
  //MpesaSTKPushRequest,
  MpesaPaymentResponse,
  MpesaCallbackRequest,
  GeneralResponse
} from '../types/api'

export class PaymentService {
  /**
   * Initiate M-Pesa STK Push
   */
  async initiateSTKPush(paymentData: MpesaPaymentRequest): Promise<GeneralResponse<MpesaPaymentResponse>> {
    const response = await apiClient.post('/v1/api/mpesa/stkpush', paymentData)
    return response.data
  }

  /**
    * Process M-Pesa callback
    */
   async processCallback(callbackData: MpesaCallbackRequest): Promise<GeneralResponse<void>> {
     const response = await apiClient.post('/v1/api/mpesa/callback', callbackData)
     return response.data
   }

  /**
   * Get payment status
   */
  async getPaymentStatus(orderId: number): Promise<GeneralResponse<{
    status: 'pending' | 'success' | 'failed'
    transaction_id?: string
  }>> {
    const response = await apiClient.get(`/payments/${orderId}/status`)
    return response.data
  }

  /**
   * Validate phone number format
   */
  validatePhoneNumber(phone: string): boolean {
    const phoneRegex = /^254[0-9]{9}$/
    return phoneRegex.test(phone.replace(/[^0-9]/g, ''))
  }

  /**
   * Format phone number to standard format
   */
  formatPhoneNumber(phone: string): string {
    const cleaned = phone.replace(/[^0-9]/g, '')
    if (cleaned.startsWith('0')) {
      return '254' + cleaned.substring(1)
    }
    return cleaned.startsWith('254') ? cleaned : '254' + cleaned
  }

  /**
   * Generate M-Pesa password
   */
  generatePassword(shortcode: string, timestamp: string): string {
    const passkey = process.env.VUE_APP_MPESA_PASSKEY || 'your_passkey'
    const password = shortcode + passkey + timestamp
    return Buffer.from(password).toString('base64')
  }

  /**
   * Generate timestamp in required format
   */
  generateTimestamp(): string {
    return new Date().toISOString().replace(/[^0-9]/g, '').slice(0, -3)
  }
}

export const paymentService = new PaymentService()