import axios, { AxiosRequestConfig } from "axios";

type Endpoint = string;

async function performGetRequest<V>(endpoint: Endpoint, config?: AxiosRequestConfig) {
    try {
        const res = await axios.get<V>("http://localhost:8000" + endpoint, {
            withCredentials: true,
            ...config
        });
        return res.data;
    } catch (e) {
        throw e;
    }
}

async function performPostRequest<V>(endpoint: Endpoint, data?: { [key: string]: any }, config?: AxiosRequestConfig) {
    try {
        const res = await axios.post<V>("http://localhost:8000" + endpoint, data, { withCredentials: true, ...config });
        return res.data;
    } catch (e) {
        throw e;
    }
}

async function performPatchRequest<V>(endpoint: Endpoint, data?: { [key: string]: any }, config?: AxiosRequestConfig) {
    try {
        const res = await axios.patch<V>("http://localhost:8000" + endpoint, data, { withCredentials: true, ...config });
        return res.data;
    } catch (e) {
        throw e;
    }
}

async function performDeleteRequest<V>(endpoint: Endpoint, config?: AxiosRequestConfig) {
    try {
        const res = await axios.delete<V>("http://localhost:8000" + endpoint, {
            withCredentials: true,
            ...config
        });
        return res.data;
    } catch (e) {
        throw e;
    }
}

const Client = {
    get: performGetRequest,
    post: performPostRequest,
    delete: performDeleteRequest,
    patch: performPatchRequest
};

export default Client;