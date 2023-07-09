/** A User in the authentication system */
export interface User {
    id: string;
    email: string;
    displayName: string;
    photoUrl: string;
    isAdmin: boolean;
    pronouns?: string;
    meetingLink?: string;
    phoneNumber?: string;
}

export interface BackendUser {
    id: string;
    email: string;
    name: string;
    picture: string;
    attributes: Attributes;
    role: string;
}

type Attributes = {
    [key: string]: any
}