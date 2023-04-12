import client from "./client";
import User from "./user";
import decode from 'jwt-decode';

enum Endpoints {
    USERS = "/users",
}

function getCookie(name: string): string | undefined {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) {
        return parts.pop()?.split(';').shift();
    }
}

function getUserIdFromToken(): string {
    const token = getCookie("token");
    if (token === undefined) {
        throw Error("No token found");
    }
    const decoded = decode<any>(token);
    const user = decoded.user;
    if (user === undefined) {
        throw Error("No user found in token");
    }
    return user.id;
}

async function getUserById(id: string): Promise<User> {
    try {
        return await client.get(`${Endpoints.USERS}/${id}`);
    } catch (e) {
        throw e;
    }
}

/**
 * Fetches profile information corresponding to the currently logged in user.
 */
function getCurrentUser(): Promise<User> {
    try {
        const id = getUserIdFromToken();
        return getUserById(id);
    } catch (e) {
        throw e;
    }
}

/**
 * Redirects the user to a Google sign in page, then creates a session with the SMU API.
 */
async function logIn() {
    try {
        return await client.post("/auth/google/login");
    } catch (e) {
        throw e;
    }
}

/**
 * Signs out the current user by removing the session cookie.
 */
async function logOut(): Promise<void> {
    try {
        return await client.post("/auth/logout");
    } catch (e) {
        throw e;
    }
}

const Auth = {
    getCurrentUser,
    logIn,
    logOut,
};


export default Auth;