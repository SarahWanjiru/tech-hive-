import { apiClient } from '../api/client.js'

export const paymentService = {
  /**
   * Initiate M-Pesa STK Push payment
   * @param {Object} paymentData - Payment data
   * @returns {Promise} API response
   */
  async initiateSTKPush(paymentData) {
    try {
      const response = await apiClient.post('/v1/api/mpesa/stkpush', {
        order_id: paymentData.orderId,
        phone_number: paymentData.phoneNumber,
        amount: parseFloat(paymentData.amount)
      })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Payment initiation failed')
    }
  },

  /**
   * Get payment status
   * @param {string|number} paymentId - Payment ID
   * @returns {Promise} API response
   */
  async getPaymentStatus(paymentId) {
    try {
      const response = await apiClient.get(`/v1/api/payments/${paymentId}`)
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to get payment status')
    }
  },

  /**
   * Get user's payment history
   * @returns {Promise} API response
   */
  async getPaymentHistory() {
    try {
      const response = await apiClient.get('/v1/api/payments/history')
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to fetch payment history')
    }
  },

  /**
   * Process payment callback (internal use)
   * @param {Object} callbackData - M-Pesa callback data
   * @returns {Promise} API response
   */
  async processCallback(callbackData) {
    try {
      const response = await apiClient.post('/v1/api/mpesa/callback', callbackData)
      return response
    } catch (error) {
      console.error('Payment callback processing failed:', error)
      throw error
    }
  },

  /**
   * Validate phone number format
   * @param {string} phoneNumber - Phone number to validate
   * @returns {boolean} Validation result
   */
  validatePhoneNumber(phoneNumber) {
    // Kenyan phone number format: 254XXXXXXXXX or 0XXXXXXXXX
    const kenyaPhoneRegex = /^(?:254|\+254|0)([7][0-9]{8})$/
    return kenyaPhoneRegex.test(phoneNumber)
  },

  /**
   * Format phone number for M-Pesa
   * @param {string} phoneNumber - Raw phone number
   * @returns {string} Formatted phone number
   */
  formatPhoneNumber(phoneNumber) {
    let cleaned = phoneNumber.replace(/\D/g, '')

    // Add 254 prefix if it starts with 0
    if (cleaned.startsWith('0')) {
      cleaned = '254' + cleaned.substring(1)
    }

    // Add + prefix if not present
    if (!cleaned.startsWith('+')) {
      cleaned = '+' + cleaned
    }

    return cleaned
  },

  /**
   * Get M-Pesa payment status text
   * @param {string} status - Payment status code
   * @returns {string} Human readable status
   */
  getPaymentStatusText(status) {
    const statusMap = {
      'pending': 'Payment Pending',
      'processing': 'Processing Payment',
      'completed': 'Payment Successful',
      'failed': 'Payment Failed',
      'cancelled': 'Payment Cancelled',
      'refunded': 'Payment Refunded'
    }
    return statusMap[status] || 'Unknown Status'
  }
}