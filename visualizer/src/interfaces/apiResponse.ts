import type { Email } from "./email";

export interface ApiErrorResponse {
    status: number;
    error: string
}

export interface EmailsSearchResponse {
    emails: Email[]
}