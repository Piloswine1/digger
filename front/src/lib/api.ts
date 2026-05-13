declare global {
  interface Window {
    __APP_CONFIG__?: { apiBaseUrl?: string }
  }
}

export const apiBaseUrl = window.__APP_CONFIG__?.apiBaseUrl ?? ''
