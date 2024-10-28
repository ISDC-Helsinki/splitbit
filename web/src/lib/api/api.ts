import createClient from "openapi-fetch";
import type { paths } from "./schema";

const api = createClient<paths>({ baseUrl: "http://localhost:8080/", credentials: "include"});

export default api
