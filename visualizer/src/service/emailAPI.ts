import axios from 'axios';

const emailAPI = axios.create({
    baseURL: 'http://localhost:8000',
});

export default emailAPI;