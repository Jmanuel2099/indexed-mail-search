import { defineStore } from 'pinia'
import axios, { AxiosError } from 'axios';

import type { Email } from '../interfaces/email'
import type { ApiErrorResponse, EmailsSearchResponse } from '@/interfaces/apiResponse';
import { emailAPI } from '@/service';

interface EmailsState {
    isLoadingEmails: boolean;
    emails: Email[];
}

export const useEmailStore = defineStore('emails', {
    state: (): EmailsState => ({
        isLoadingEmails: false,
        emails: [],
    }),
    getters: {

    },
    actions: {
        setisLoadingEmails() {
            this.isLoadingEmails = true;
        },
        setEmails(emails: Email[]) {
            this.emails = emails;
            this.isLoadingEmails = false;
        },
        async searchEmailsByTerm(termQuery: string) {
            if (termQuery.length === 0) {
                this.setEmails([]);
            }

            this.setisLoadingEmails()
            try {
                const emailsResponse = await emailAPI.get<EmailsSearchResponse>(`/search?term=${termQuery}`);
                this.setEmails(emailsResponse.data.emails)

                console.log(emailsResponse.data.emails)
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    const apiErr = error as AxiosError<ApiErrorResponse>
                    console.error('API Error -', apiErr.response?.data.error)
                } else {
                    console.error('Unknown Error -', error)
                }

                this.setEmails([]);
            }
        },
    }
});