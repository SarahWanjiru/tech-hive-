import { ref, reactive, onMounted, onUnmounted } from 'vue'

export function useTouchGestures(elementRef) {
  const touchStart = ref(null)
  const touchCurrent = ref(null)
  const isGesturing = ref(false)

  const gestureState = reactive({
    direction: null, // 'left', 'right', 'up', 'down'
    distance: 0,
    velocity: 0,
    startTime: 0,
    isSwipe: false,
    isLongPress: false
  })

  let longPressTimer = null
  let lastMoveTime = 0

  const MIN_SWIPE_DISTANCE = 50
  const MIN_SWIPE_VELOCITY = 0.3
  const LONG_PRESS_DELAY = 500

  const calculateDistance = (start, end) => {
    const dx = end.clientX - start.clientX
    const dy = end.clientY - start.clientY
    return Math.sqrt(dx * dx + dy * dy)
  }

  const calculateDirection = (start, end) => {
    const dx = end.clientX - start.clientX
    const dy = end.clientY - start.clientY
    const absDx = Math.abs(dx)
    const absDy = Math.abs(dy)

    if (Math.max(absDx, absDy) < 10) return null

    if (absDx > absDy) {
      return dx > 0 ? 'right' : 'left'
    } else {
      return dy > 0 ? 'down' : 'up'
    }
  }

  const calculateVelocity = (start, end, startTime, endTime) => {
    const distance = calculateDistance(start, end)
    const time = endTime - startTime
    return time > 0 ? distance / time : 0
  }

  const handleTouchStart = (e) => {
    const touch = e.touches[0]
    touchStart.value = {
      x: touch.clientX,
      y: touch.clientY,
      time: Date.now()
    }
    touchCurrent.value = { ...touchStart.value }
    gestureState.startTime = Date.now()
    isGesturing.value = true

    // Long press detection
    longPressTimer = setTimeout(() => {
      gestureState.isLongPress = true
      elementRef.value?.dispatchEvent(new CustomEvent('long-press', {
        detail: { x: touch.clientX, y: touch.clientY }
      }))
    }, LONG_PRESS_DELAY)
  }

  const handleTouchMove = (e) => {
    if (!touchStart.value || !isGesturing.value) return

    const touch = e.touches[0]
    touchCurrent.value = {
      x: touch.clientX,
      y: touch.clientY,
      time: Date.now()
    }

    const now = Date.now()
    const dx = touchCurrent.value.x - touchStart.value.x
    const dy = touchCurrent.value.y - touchStart.value.y
    const distance = Math.sqrt(dx * dx + dy * dy)

    gestureState.distance = distance
    gestureState.direction = calculateDirection(touchStart.value, touchCurrent.value)
    gestureState.velocity = calculateVelocity(
      touchStart.value,
      touchCurrent.value,
      gestureState.startTime,
      now
    )

    // Cancel long press if moved too much
    if (distance > 10) {
      clearTimeout(longPressTimer)
      gestureState.isLongPress = false
    }

    // Emit move events
    elementRef.value?.dispatchEvent(new CustomEvent('touch-move', {
      detail: { ...gestureState, current: touchCurrent.value }
    }))
  }

  const handleTouchEnd = (e) => {
    if (!touchStart.value || !isGesturing.value) return

    clearTimeout(longPressTimer)
    const touch = e.changedTouches[0]
    const endTouch = {
      x: touch.clientX,
      y: touch.clientY,
      time: Date.now()
    }

    const distance = calculateDistance(touchStart.value, endTouch)
    const velocity = calculateVelocity(touchStart.value, endTouch, gestureState.startTime, endTouch.time)
    const direction = calculateDirection(touchStart.value, endTouch)

    // Determine if it's a swipe
    const isSwipe = distance >= MIN_SWIPE_DISTANCE && velocity >= MIN_SWIPE_VELOCITY

    gestureState.isSwipe = isSwipe
    gestureState.direction = direction
    gestureState.distance = distance
    gestureState.velocity = velocity

    // Emit swipe or tap events
    if (isSwipe) {
      elementRef.value?.dispatchEvent(new CustomEvent('swipe', {
        detail: { ...gestureState, end: endTouch }
      }))
    } else {
      elementRef.value?.dispatchEvent(new CustomEvent('tap', {
        detail: { position: endTouch, duration: endTouch.time - gestureState.startTime }
      }))
    }

    // Reset state
    touchStart.value = null
    touchCurrent.value = null
    isGesturing.value = false
    gestureState.isLongPress = false
  }

  const setupElementListeners = (element) => {
    if (!element) return

    element.addEventListener('touchstart', handleTouchStart, { passive: false })
    element.addEventListener('touchmove', handleTouchMove, { passive: false })
    element.addEventListener('touchend', handleTouchEnd, { passive: false })
  }

  const removeElementListeners = (element) => {
    if (!element) return

    element.removeEventListener('touchstart', handleTouchStart)
    element.removeEventListener('touchmove', handleTouchMove)
    element.removeEventListener('touchend', handleTouchEnd)
  }

  onMounted(() => {
    if (elementRef.value) {
      setupElementListeners(elementRef.value)
    }
  })

  onUnmounted(() => {
    if (elementRef.value) {
      removeElementListeners(elementRef.value)
    }
    if (longPressTimer) {
      clearTimeout(longPressTimer)
    }
  })

  return {
    gestureState,
    isGesturing,
    setupElementListeners,
    removeElementListeners
  }
}

export function useSwipeNavigation(items, currentIndex = ref(0)) {
  const navigateNext = () => {
    if (currentIndex.value < items.length - 1) {
      currentIndex.value++
      return items[currentIndex.value]
    }
    return null
  }

  const navigatePrev = () => {
    if (currentIndex.value > 0) {
      currentIndex.value--
      return items[currentIndex.value]
    }
    return null
  }

  const goToIndex = (index) => {
    if (index >= 0 && index < items.length) {
      currentIndex.value = index
      return items[index]
    }
    return null
  }

  return {
    currentIndex,
    navigateNext,
    navigatePrev,
    goToIndex,
    currentItem: computed(() => items[currentIndex.value] || null)
  }
}

export function usePullToRefresh(onRefresh) {
  const isPulling = ref(false)
  const pullDistance = ref(0)
  const startY = ref(0)
  const threshold = 80

  const handleTouchStart = (e) => {
    if (window.scrollY === 0) {
      startY.value = e.touches[0].clientY
      isPulling.value = true
    }
  }

  const handleTouchMove = (e) => {
    if (!isPulling.value || window.scrollY > 0) return

    const currentY = e.touches[0].clientY
    const distance = currentY - startY.value

    if (distance > 0) {
      pullDistance.value = Math.min(distance * 0.5, threshold + 40)
      e.preventDefault()
    }
  }

  const handleTouchEnd = () => {
    if (!isPulling.value) return

    if (pullDistance.value > threshold) {
      onRefresh()
    }

    isPulling.value = false
    pullDistance.value = 0
  }

  return {
    isPulling,
    pullDistance,
    handleTouchStart,
    handleTouchMove,
    handleTouchEnd
  }
}