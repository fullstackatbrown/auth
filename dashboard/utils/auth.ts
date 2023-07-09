import { useEffect, useState } from "react";
import { User, BackendUser } from "./user"
import client from "./client";

/**
 * Redirects the user to a Google sign in page, then creates a session with the SMU API.
 */
export function login(authBaseHost: string, applicationHome: string): void {
  location.href = `${authBaseHost}/v1/auth/google/login?from=${applicationHome}`
}

/**
 * Signs out the current user by removing the session cookie.
 */
export function logout(): Promise<void> {
  try {
    return client.get("/v1/auth/logout");
  } catch (e) {
    return Promise.reject(e);
  }
}

/**
 * React hook for getting the current user.
 */
export function useUser() {
  const [user, setUser] = useState<User>();
  const [error, setError] = useState<Error>();

  useEffect(() => {
    getUser()
      .then((user) => setUser(user))
      .catch((e) => setError(e));
  }, []);

  return {
    user,
    error,
    loading: !user && !error,
  };
}

async function getUser(): Promise<User> {
  try {
    const user = await client.get<BackendUser>(`/v1/auth/user`)
    return castUser(user)
  } catch (e) {
    return Promise.reject(e);
  }
}

function castUser(user: BackendUser): User {
  return { id: user.id, email: user.email, displayName: user.name, photoUrl: user.picture, isAdmin: false }
}