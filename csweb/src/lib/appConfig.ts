export interface AppConfig {
	graphqlEndpoint: string;
	graphqlWSEndpoint: string;
	fileEndpoint: string;
	defaultResultsPerPage: number;

	tokenRefreshMinutes: number;

	loginUrl: string;
	signoutUrl: string;
	refreshUrl: string;
}

//---TODO: Load from file or env
export const appConfig: AppConfig = {
	graphqlEndpoint: 'http://localhost:5000/query',
	graphqlWSEndpoint: 'ws://localhost:5000/query',
	fileEndpoint: 'http://localhost:5000/files/',
	defaultResultsPerPage: 20,

	tokenRefreshMinutes: 4,

	loginUrl: 'http://localhost:5000/login',
	signoutUrl: 'http://localhost:5000/signout',
	refreshUrl: 'http://localhost:5000/refresh'
};
