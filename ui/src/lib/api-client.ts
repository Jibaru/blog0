export interface ErrorResp {
  error: string;
}

export interface CreateCommentReq {
  body: string;
  parent_id?: string;
}

export interface CreatePostReq {
  raw_markdown: string;
  slug: string;
  title: string;
  publish?: boolean;
}

export interface UpdatePostReq {
  publish?: boolean;
  raw_markdown?: string;
  slug?: string;
  title?: string;
}

export interface AuthorInfo {
  id: string;
  name: string;
}

export interface BookmarkPostResp {
  bookmark_id: string;
  bookmarked: boolean;
  created_at: string;
  post_slug: string;
}

export interface CommentInfo {
  author: AuthorInfo;
  body: string;
  created_at: string;
  id: string;
  parent_id?: string;
}

export interface CreateCommentResp {
  author: AuthorInfo;
  body: string;
  created_at: string;
  id: string;
  parent_id?: string;
  post_slug: string;
}

export interface CreatePostResp {
  author_id: string;
  created_at: string;
  id: string;
  published_at?: string;
  raw_markdown: string;
  slug: string;
  title: string;
  updated_at: string;
}

export interface DeletePostResp {
  message: string;
  success: boolean;
}

export interface FollowUserResp {
  followers_count: number;
  following: boolean;
}

export interface TopPostInfo {
  id: string;
  likes_count: number;
  slug: string;
  title: string;
}

export interface GetAuthorInfoResp {
  followers_count: number;
  id: string;
  name: string;
  posts_count: number;
  top_posts: TopPostInfo[];
}

export interface GetPostBySlugResp {
  author: AuthorInfo;
  comments: CommentInfo[];
  id: string;
  likes_count: number;
  published_at?: string;
  raw_markdown: string;
  slug: string;
  title: string;
}

export interface MyPostItem {
  created_at: string;
  id: string;
  published_at?: string;
  slug: string;
  status: string;
  title: string;
  updated_at: string;
}

export interface ListMyPostsResp {
  items: MyPostItem[];
  page: number;
  per_page: number;
  total: number;
}

export interface PostItem {
  author: string;
  author_id: string;
  published_at?: string;
  slug: string;
  title: string;
  comment_count: number;
  like_count: number;
}

export interface ListPostsResp {
  items: PostItem[];
  page: number;
  per_page: number;
  total: number;
}

export interface ToggleLikeResp {
  liked: boolean;
  likes_count: number;
}

export interface UnbookmarkPostResp {
  bookmarked: boolean;
}

export interface UnfollowUserResp {
  followers_count: number;
  following: boolean;
}

export interface ProfilePost {
  id: string;
  title: string;
}

export interface ProfileUser {
  id: string;
  username: string;
}

export interface GetProfileResp {
  bookmarks: ProfilePost[];
  following: ProfileUser[];
  liked_posts: ProfilePost[];
}

export interface UpdatePostResp {
  author_id: string;
  created_at: string;
  id: string;
  published_at?: string;
  raw_markdown: string;
  slug: string;
  title: string;
  updated_at: string;
}

export interface ListPostsParams {
  page?: number;
  per_page?: number;
  order?: string;
  [key: string]: string | number | boolean | undefined;
}

export interface ListMyPostsParams {
  page?: number;
  per_page?: number;
  order?: 'ASC' | 'DESC';
  [key: string]: string | number | boolean | undefined;
}

export interface ApiClientConfig {
  baseUrl?: string;
  apiToken?: string;
  defaultHeaders?: Record<string, string>;
}

export class ApiError extends Error {
  constructor(
    public status: number,
    public response: ErrorResp | string,
    message?: string
  ) {
    super(message || (typeof response === 'string' ? response : response.error));
    this.name = 'ApiError';
  }
}

export class Blog0ApiClient {
  private baseUrl: string;
  private apiToken?: string;
  private defaultHeaders: Record<string, string>;

  constructor(config: ApiClientConfig = {}) {
    this.baseUrl = config.baseUrl || 'https://blog0-backend.vercel.app';
    this.apiToken = config.apiToken;
    this.defaultHeaders = {
      'Content-Type': 'application/json',
      ...config.defaultHeaders,
    };
  }

  setApiToken(token: string): void {
    this.apiToken = token;
  }

  private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    const url = `${this.baseUrl}/api/v1${endpoint}`;
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

    if (response.status === 204 || response.headers.get('content-length') === '0') {
      return {} as T;
    }

    return response.json();
  }

  private buildQueryString(params: Record<string, string | number | boolean | undefined>): string {
    const searchParams = new URLSearchParams();
    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined && value !== null) {
        searchParams.append(key, value.toString());
      }
    });
    return searchParams.toString();
  }

  async startOAuth(provider: string): Promise<void> {
    // Redirect directly to OAuth URL
    window.location.href = `${this.baseUrl}/api/v1/auth/${provider}`;
  }

  async oauthCallback(provider: string): Promise<void> {
    // Redirect directly to callback URL with current search params
    window.location.href = `${this.baseUrl}/api/v1/auth/${provider}/callback${window.location.search}`;
  }

  async listPosts(params: ListPostsParams = {}): Promise<ListPostsResp> {
    const queryString = this.buildQueryString(params);
    const endpoint = `/posts${queryString ? `?${queryString}` : ''}`;
    return this.request<ListPostsResp>(endpoint);
  }

  async getPostBySlug(slug: string): Promise<GetPostBySlugResp> {
    return this.request<GetPostBySlugResp>(`/posts/${slug}`);
  }

  async toggleLike(slug: string): Promise<ToggleLikeResp> {
    return this.request<ToggleLikeResp>(`/posts/${slug}/likes`, {
      method: 'POST',
    });
  }

  async bookmarkPost(slug: string): Promise<BookmarkPostResp> {
    return this.request<BookmarkPostResp>(`/posts/${slug}/bookmarks`, {
      method: 'POST',
    });
  }

  async unbookmarkPost(slug: string): Promise<UnbookmarkPostResp> {
    return this.request<UnbookmarkPostResp>(`/posts/${slug}/bookmarks`, {
      method: 'DELETE',
    });
  }

  async createComment(slug: string, comment: CreateCommentReq): Promise<CreateCommentResp> {
    return this.request<CreateCommentResp>(`/posts/${slug}/comments`, {
      method: 'POST',
      body: JSON.stringify(comment),
    });
  }

  async listMyPosts(params: ListMyPostsParams = {}): Promise<ListMyPostsResp> {
    const queryString = this.buildQueryString(params);
    const endpoint = `/me/posts${queryString ? `?${queryString}` : ''}`;
    return this.request<ListMyPostsResp>(endpoint);
  }

  async createPost(post: CreatePostReq): Promise<CreatePostResp> {
    return this.request<CreatePostResp>('/me/posts', {
      method: 'POST',
      body: JSON.stringify(post),
    });
  }

  async updatePost(slug: string, post: UpdatePostReq): Promise<UpdatePostResp> {
    return this.request<UpdatePostResp>(`/me/posts/${slug}`, {
      method: 'PUT',
      body: JSON.stringify(post),
    });
  }

  async deletePost(slug: string): Promise<DeletePostResp> {
    return this.request<DeletePostResp>(`/me/posts/${slug}`, {
      method: 'DELETE',
    });
  }

  async getAuthorInfo(authorId: string): Promise<GetAuthorInfoResp> {
    return this.request<GetAuthorInfoResp>(`/users/${authorId}`);
  }

  async followUser(authorId: string): Promise<FollowUserResp> {
    return this.request<FollowUserResp>(`/users/${authorId}/follow`, {
      method: 'POST',
    });
  }

  async unfollowUser(authorId: string): Promise<UnfollowUserResp> {
    return this.request<UnfollowUserResp>(`/users/${authorId}/follow`, {
      method: 'DELETE',
    });
  }

  async getProfile(): Promise<GetProfileResp> {
    return this.request<GetProfileResp>('/me/profile');
  }
}

// Create a shared global instance
const globalApiClient = new Blog0ApiClient();

export default Blog0ApiClient;
export { globalApiClient };
