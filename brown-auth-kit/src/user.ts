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
