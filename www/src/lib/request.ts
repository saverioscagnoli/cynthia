type ApiError = {
  status: number;
  message: string;
  details?: unknown;
};

type Result<T> = { ok: true; data: T } | { ok: false; error: ApiError };

async function api<T>(
  endpoint: string,
  options?: RequestInit
): Promise<Result<T>> {
  try {
    let url = __API_URL__ + endpoint;
    let res = await fetch(url, options);

    let contentType = res.headers.get("content-type");
    let isJson = contentType?.includes("application/json");

    let data = isJson ? await res.json() : await res.text();

    if (!res.ok) {
      return {
        ok: false,
        error: {
          status: res.status,
          message: data?.message || "Request failed",
          details: data
        }
      };
    }

    return {
      ok: true,
      data: data as T
    };
  } catch (err) {
    return {
      ok: false,
      error: {
        status: 0,
        message: "Network error",
        details: err
      }
    };
  }
}

export { api };
export type { Result, ApiError };
