import { defineStore } from 'pinia'
import { paymentService } from '../services/payment.service'
import type { MpesaPaymentRequest, MpesaCallbackRequest } from '../types/api'

interface PaymentState {
  loading: boolean
  error: string | null
  paymentStatus: 'idle' | 'processing' | 'success' | 'failed'
}

export const usePaymentStore = defineStore('payment', {
  state: (): PaymentState => ({
    loading: false,
    error: null,
    paymentStatus: 'idle'
  }),

  getters: {
    isProcessing: (state: PaymentState): boolean => state.paymentStatus === 'processing',
    isSuccess: (state: PaymentState): boolean => state.paymentStatus === 'success',
    isFailed: (state: PaymentState): boolean => state.paymentStatus === 'failed'
  },

  actions: {
    /**
     * Initiate M-Pesa STK Push payment
     */
    async initiateSTKPush(paymentData: MpesaPaymentRequest): Promise<void> {
      if (!paymentService.validatePhoneNumber(paymentData.phone_number)) {
        throw new Error('Invalid phone number format')
      }

      this.loading = true
      this.error = null
      this.paymentStatus = 'processing'

      try {
        const response = await paymentService.initiateSTKPush(paymentData)

        if (response.code === 200) {
          this.paymentStatus = 'success'
        } else {
          this.paymentStatus = 'failed'
          this.error = response.message
        }
      } catch (error: any) {
        this.paymentStatus = 'failed'
        this.error = error.message || 'Payment initiation failed'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Process M-Pesa callback
     */
    async processCallback(callbackData: MpesaCallbackRequest): Promise<void> {
      try {
        await paymentService.processCallback(callbackData)

        if (callbackData.result_code === 0) {
          this.paymentStatus = 'success'
        } else {
          this.paymentStatus = 'failed'
          this.error = callbackData.result_desc
        }
      } catch (error: any) {
        this.error = error.message || 'Callback processing failed'
        throw error
      }
    },

    /**
     * Get payment status for order
     */
    async getPaymentStatus(orderId: number): Promise<string> {
      try {
        const response = await paymentService.getPaymentStatus(orderId)
        return response.data.status
      } catch (error: any) {
        console.error('Failed to get payment status:', error)
        return 'pending'
      }
    },

    /**
     * Validate phone number
     */
    validatePhoneNumber(phone: string): boolean {
      return paymentService.validatePhoneNumber(phone)
    },

    /**
     * Format phone number
     */
    formatPhoneNumber(phone: string): string {
      return paymentService.formatPhoneNumber(phone)
    },

    /**
     * Reset payment state
     */
    resetPaymentState(): void {
      this.paymentStatus = 'idle'
      this.error = null
      this.loading = false
    },

    /**
     * Clear error state
     */
    clearError(): void {
      this.error = null
    }
  }
})