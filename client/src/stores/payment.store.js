import { defineStore } from 'pinia'
import { paymentService } from '../services/payment.service.js'

export const usePaymentStore = defineStore('payment', {
  state: () => ({
    loading: false,
    error: null,
    paymentHistory: [],
    currentPayment: null
  }),

  getters: {
    hasPaymentHistory: (state) => {
      return state.paymentHistory.length > 0
    }
  },

  actions: {
    /**
     * Initiate M-Pesa STK Push payment
     * @param {Object} paymentData - Payment data
     */
    async initiateSTKPush(paymentData) {
      this.loading = true
      this.error = null

      try {
        const response = await paymentService.initiateSTKPush(paymentData)
        this.currentPayment = response.data
        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Get payment status
     * @param {string|number} paymentId - Payment ID
     */
    async getPaymentStatus(paymentId) {
      try {
        const response = await paymentService.getPaymentStatus(paymentId)
        this.currentPayment = response.data
        return response
      } catch (error) {
        this.error = error.message
        throw error
      }
    },

    /**
     * Get payment history
     */
    async fetchPaymentHistory() {
      this.loading = true
      this.error = null

      try {
        const response = await paymentService.getPaymentHistory()
        this.paymentHistory = response.data.payments || response.data || []
        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Validate phone number
     * @param {string} phoneNumber - Phone number to validate
     * @returns {boolean} Validation result
     */
    validatePhoneNumber(phoneNumber) {
      return paymentService.validatePhoneNumber(phoneNumber)
    },

    /**
     * Format phone number for M-Pesa
     * @param {string} phoneNumber - Raw phone number
     * @returns {string} Formatted phone number
     */
    formatPhoneNumber(phoneNumber) {
      return paymentService.formatPhoneNumber(phoneNumber)
    },

    /**
     * Get payment status text
     * @param {string} status - Payment status code
     * @returns {string} Human readable status
     */
    getPaymentStatusText(status) {
      return paymentService.getPaymentStatusText(status)
    },

    /**
     * Clear error state
     */
    clearError() {
      this.error = null
    },

    /**
     * Clear current payment
     */
    clearCurrentPayment() {
      this.currentPayment = null
    }
  }
})