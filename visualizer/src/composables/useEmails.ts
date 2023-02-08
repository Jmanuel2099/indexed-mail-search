import { useEmailStore } from "@/stores/emial"
import { storeToRefs } from "pinia"

export const useEmails = () => {

    const emailsStore = useEmailStore();
    const { emails, isLoadingEmails } = storeToRefs(emailsStore);

    const searchEmailsByTerm = (temr: string) => {
        emailsStore.searchEmailsByTerm(temr);
    };

    return {
        // State
        emails,
        isLoadingEmails,
        //Getters
        //Actions
        searchEmailsByTerm
    };
};