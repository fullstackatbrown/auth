import client from "./client";
import { User } from "./user";
import decode from "jwt-decode";
import { useEffect, useState } from "react";

enum Endpoints {
  USERS = "/users",
}

function getCookie(name: string): string | undefined {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) {
    return parts.pop()?.split(";").shift();
  }
}

function getUserIdFromToken(tokenName: string): string {
  const token = getCookie(tokenName);
  if (token === undefined) {
    throw Error("No token found");
  }
  const decoded = decode<any>(token);
  const user = decoded.user as unknown;
  if (
    !user ||
    typeof user !== "object" ||
    !("id" in user) ||
    typeof user.id !== "string"
  ) {
    throw Error("No user found in token");
  }
  return user.id;
}

export function getUserById(id: string): Promise<User> {
  return client.get(`${Endpoints.USERS}/${id}`);
}

/**
 * Fetches profile information corresponding to the currently logged in user.
 */
export function getCurrentUser(): Promise<User> {
  try {
    const id = getUserIdFromToken("fsab-session");
    return getUserById(id);
  } catch (e) {
    return Promise.reject(e);
  }
}

/**
 * Redirects the user to a Google sign in page, then creates a session with the SMU API.
 */
export function signIn(authBaseHost: string, applicationHome: string): void {
  location.href = `${authBaseHost}/login?from=${applicationHome}`
}

/**
 * Signs out the current user by removing the session cookie.
 */
export function signOut(): Promise<void> {
  return client.post("/auth/logout");
}

/**
 * React hook for getting the current user.
 */
export function useUser() {
  const [user, setUser] = useState<User>();
  const [error, setError] = useState<Error>();

  useEffect(() => {
    getCurrentUser()
      .then((user) => setUser(user))
      .catch((e) => setError(e));
  }, []);

  return {
    user,
    error,
    loading: !user && !error,
  };
}
