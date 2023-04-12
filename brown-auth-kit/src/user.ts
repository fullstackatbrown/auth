/** A User in the authentication system */
export default interface User {
    id: string;
    email: string;
    displayName: string;
    photoUrl: string;
    isAdmin: boolean;
    pronouns?: string;
    meetingLink?: string;
    phoneNumber?: string;
}