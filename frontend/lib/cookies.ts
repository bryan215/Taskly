import Cookies from 'js-cookie';

const TOKEN_COOKIE = 'token';

interface TokenPayload {
  user_id: number;
  username: string;
  exp: number;
}

function decodeToken(token: string): TokenPayload | null {
  try {
    const base64Url = token.split('.')[1];
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split('')
        .map((c) => '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2))
        .join('')
    );
    return JSON.parse(jsonPayload);
  } catch {
    return null;
  }
}

export const cookieUtils = {
  setToken(token: string) {
    Cookies.set(TOKEN_COOKIE, token, { expires: 7 });
  },

  getToken(): string | null {
    return Cookies.get(TOKEN_COOKIE) || null;
  },

  getUserId(): number | null {
    const token = this.getToken();
    if (!token) return null;
    const payload = decodeToken(token);
    return payload?.user_id || null;
  },

  getUsername(): string | null {
    const token = this.getToken();
    if (!token) return null;
    const payload = decodeToken(token);
    return payload?.username || null;
  },

  isTokenExpired(): boolean {
    const token = this.getToken();
    if (!token) return true;
    const payload = decodeToken(token);
    if (!payload || !payload.exp) return true;
    return payload.exp * 1000 < Date.now();
  },

  clearAuth() {
    Cookies.remove(TOKEN_COOKIE);
  },
};





