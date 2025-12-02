import Cookies from 'js-cookie';

const USER_ID_COOKIE = 'user_id';
const USERNAME_COOKIE = 'username';
const TOKEN_COOKIE = 'token';

export const cookieUtils = {
  setUserId(userId: number) {
    Cookies.set(USER_ID_COOKIE, userId.toString(), { expires: 7 }); // 7 días
  },

  getUserId(): number | null {
    const userId = Cookies.get(USER_ID_COOKIE);
    return userId ? parseInt(userId, 10) : null;
  },

  setUsername(username: string) {
    Cookies.set(USERNAME_COOKIE, username, { expires: 7 });
  },

  getUsername(): string | null {
    return Cookies.get(USERNAME_COOKIE) || null;
  },

  setToken(token: string) {
    Cookies.set(TOKEN_COOKIE, token, { expires: 7 });
  },

  getToken(): string | null {
    return Cookies.get(TOKEN_COOKIE) || null;
  },

  clearAuth() {
    Cookies.remove(USER_ID_COOKIE);
    Cookies.remove(USERNAME_COOKIE);
    Cookies.remove(TOKEN_COOKIE);
  },
};





