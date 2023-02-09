import type { Email } from "@/interfaces/email";
import { useEmailStore } from "@/stores/emial"
import { storeToRefs } from "pinia"

export const useEmails = () => {

    const emailsStore = useEmailStore();
    const { emails, isLoadingEmails, selectedEmail, selectedEmailReady } = storeToRefs(emailsStore);

    const searchEmailsByTerm = (temr: string) => {
        emailsStore.searchEmailsByTerm(temr);
    };

    const onSelectEmail =(email: Email) =>{
        emailsStore.setSelectedEmail(email)
    };
    return {
        // State
        emails,
        isLoadingEmails,
        selectedEmail,
        //Getters
        selectedEmailReady,
        //Actions
        searchEmailsByTerm,
        onSelectEmail

    };
};