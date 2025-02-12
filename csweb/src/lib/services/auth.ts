import { appConfig } from '$lib/appConfig';
import { type User, CurrentUserDocument } from '$lib/graphql/generated/sdk';
import { getCookie, deleteCookie } from '$lib/utils/helpers';
import { getApolloClient } from '$lib/graphql/gqlclient';

let user: User | undefined;
let refreshId: ReturnType<typeof setInterval>;

type UpdateLogin = {
	email: string;
	password: string;
};

export function authService() {
	/**
	 * Checks the credentials against the auth store and if successful, returns the
	 * required tokens
	 * @param creds the credentials for a login which contains an email and password
	 * @returns
	 */
	function login(creds: UpdateLogin): Promise<boolean> {
		const request = new Request(appConfig.loginUrl, {
			method: 'POST',
			body: JSON.stringify({ username: creds.email, password: creds.password }),
			headers: {
				'Content-Type': 'application/json'
			}
		});

		return fetch(request)
			.then((response) => response.json())
			.then((data) => {
				if (data.success) {
					//---TODO: come up with a better strategy for local handling
					storeTokens(data.accessToken, data.refreshToken);

					decodeUser(data.accessToken);

					//---refresh the token on the configured interval
					refreshCycle();
					console.log('setting refresh interval', refreshId);

					return true;
				} else {
					return false;
				}
			})
			.catch((err) => {
				console.error('login error', err);

				return false;
			});
	}

	/**
	 * calls the signout endpoint and cleans up browser objects if successful
	 */
	function signout(): Promise<boolean> {
		const refreshToken = getRefreshToken();
		const accessToken = getAccessToken();

		const request = new Request(appConfig.signoutUrl, {
			method: 'POST',
			body: JSON.stringify({ refreshToken }),
			headers: {
				'Content-Type': 'application/json',
				authorization: accessToken
			}
		});

		return fetch(request)
			.then((response) => {
				if (response.status === 403) {
					//---access token already expired
					cleanup();
					return true;
				}
				response.json();
			})
			.then(() => {
				cleanup();
				return true;
			})
			.catch((err) => {
				console.error('signout error', err);
				cleanup();
				return true;
			});
	}

	/**
	 * remove all traces of the user's login footprint
	 */
	function cleanup() {
		deleteCookie('accessToken');
		deleteCookie('refreshToken');

		clearInterval(refreshId);
	}

	/**
	 * calls the refresh endpoint. if successful, replaces
	 * the access token with a new one, extending the expiration date
	 */
	function refresh(): Promise<boolean> {
		const refreshToken = getRefreshToken();
		console.log('refreshing token...');

		const request = new Request(appConfig.refreshUrl, {
			method: 'POST',
			body: JSON.stringify({ refreshToken }),
			headers: {
				'Content-Type': 'application/json',
				authorization: getAccessToken()
			}
		});

		return fetch(request)
			.then((response) => response.json())
			.then((data) => {
				if (data.success) {
					//---TODO: come up with a better strategy for local handling
					storeTokens(data.accessToken, data.refreshToken);

					console.log('token refreshed...');

					return true;
				} else {
					return false;
				}
			})
			.catch((err) => {
				console.error('refresh error', err);

				return false;
			});
	}

	/**
	 * This will start an inverval that updates the access token behind the scenes.
	 * It is
	 */
	function refreshCycle() {
		const accessToken = getAccessToken();
		const refreshToken = getRefreshToken();

		if (accessToken && refreshToken) {
			console.log('begin refresh cycle');

			//---clear current refresh cycle if exists
			clearInterval(refreshId);

			//---begin a new refresh cycle
			refreshId = setInterval(refresh, appConfig.tokenRefreshMinutes * 1000 * 60);
		}
	}

	/**
	 * persist the access and refresh tokens.  This is an abstraction as the
	 * logic will likely change
	 * @param access the access token
	 * @param refresh the refresh token
	 */
	function storeTokens(access: string, refresh: string) {
		document.cookie = `accessToken=${access}; path=/`;
		document.cookie = `refreshToken=${refresh}; path=/`;
	}

	/**
	 * Get the currently logged in user
	 * @returns the currently logged in user
	 */
	function currentUser(): User | undefined {
		user = decodeUser(getAccessToken());
		if (user) {
			return user;
		}

		return undefined;
	}

	/**
	 * Tries to get a current user to ensure login is active
	 */
	function pingUser(): Promise<User> {
		const client = getApolloClient();

		return client
			.query({ query: CurrentUserDocument })
			.then((u) => {
				return u.data.currentUser;
			})
			.catch(() => {
				return null;
			});
	}

	/**
	 * Check if the user is logged in and return true if so.  Otherwise, it redirects to login page
	 * @returns true if the user is logged in.
	 */
	function authCheck(): Promise<boolean> {
		const accessToken = getAccessToken();
		if (accessToken) {
			return pingUser().then((r) => {
				if (r != null) {
					return true;
				} else {
					return false;
				}
			});
		}

		return new Promise<boolean>((resolve) => {
			resolve(false);
		});
	}

	/**
	 * Return authentication headers based on the current login status
	 * @returns a headers object
	 */
	function getAuthHeaders() {
		const accessToken = getAccessToken();

		if (accessToken) {
			return {
				authorization: accessToken
			};
		}

		return { authorization: '' };
	}

	/**
	 * return the current access token of the logged in user
	 * @returns the current access token if available
	 */
	function getAccessToken(): string {
		return getCookie('accessToken');
	}

	/**
	 * return the current refresh token of the logged in user
	 * @returns the current refresh token if available
	 */
	function getRefreshToken() {
		return getCookie('refreshToken');
	}

	/**
	 * Take an access token and extract the user properties
	 * @param accessToken - the access token for the current login
	 * @returns a populated user object if the token is valid.  Undefined otherwise
	 */
	function decodeUser(accessToken: string): User | undefined {
		if (!accessToken) {
			return undefined;
		}

		const decoded = JSON.parse(atob(accessToken.split('.')[1]));
		console.log(decoded)

		user = {
			id: decoded.sub,
			firstName: decoded.given_name,
			lastName: decoded.family_name,
			email: decoded.email,
			profileImage: decoded.profileImage
		};

		return user as User;
	}

	/**
	 * helper function to output state of the auth service
	 * Should be deleted for security purposes
	 */
	function showData() {
		console.log('--------- Auth Service Status ---------');
		console.log('accessToken', getAccessToken());
		console.log('refreshToken', getRefreshToken());
		console.log('user', user);
		console.log('---------------------------------------');
	}

	return {
		login,
		signout,
		currentUser,
		authCheck,
		getAuthHeaders,
		getAccessToken,
		refreshCycle
	};
}
