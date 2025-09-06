export interface ErrorResp {
  error: string;
}

export interface CreatePostReq {
  raw_markdown: string;
  slug: string;
  title: string;
  publish?: boolean;
}

export interface CreatePostResp {
  author_id: string;
  created_at: string;
  id: string;
  published_at?: string;
  raw_markdown: string;
  slug: string;
  summary: string;
  tags: string[];
  title: string;
  updated_at: string;
}

export interface ApiClientConfig {
  baseUrl?: string;
  apiToken?: string;
  defaultHeaders?: Record<string, string>;
  onTokenExpired?: () => void;
}

export class ApiError extends Error {
  constructor(
    public status: number,
    public response: ErrorResp | string,
    message?: string
  ) {
    super(
      message || (typeof response === "string" ? response : response.error)
    );
    this.name = "ApiError";
  }
}

export class Blog0ApiClient {
  private baseUrl: string;
  private apiToken?: string;
  private defaultHeaders: Record<string, string>;

  constructor(config: ApiClientConfig = {}) {
    this.baseUrl = config.baseUrl || "https://blog0-backend.vercel.app";
    this.apiToken = config.apiToken;
    this.defaultHeaders = {
      "Content-Type": "application/json",
      ...config.defaultHeaders,
    };
  }

  setApiToken(token: string): void {
    this.apiToken = token;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseUrl}/api/p/v1${endpoint}`;
    const headers: Record<string, string> = {
      ...this.defaultHeaders,
      ...(options.headers as Record<string, string>),
    };

    if (this.apiToken) {
      headers.Authorization = `${this.apiToken}`;
    }

    const response = await fetch(url, {
      ...options,
      headers,
    });

    if (!response.ok) {
      let errorResponse: ErrorResp | string;
      try {
        errorResponse = await response.json();
      } catch {
        errorResponse = await response.text();
      }
      throw new ApiError(response.status, errorResponse);
    }

    return response.json();
  }

  async createPost(post: CreatePostReq): Promise<CreatePostResp> {
    return this.request<CreatePostResp>("/posts", {
      method: "POST",
      body: JSON.stringify(post),
    });
  }
}

// Instancia global opcional
const globalApiClient = new Blog0ApiClient();

export default Blog0ApiClient;
export { globalApiClient };
