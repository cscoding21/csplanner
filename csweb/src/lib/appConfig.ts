export interface AppConfig {
    graphqlEndpoint: string; 
    graphqlWSEndpoint: string;
    fileEndpoint: string;
    defaultResultsPerPage: number;
  }
    
  //---TODO: Load from file or env
  export const appConfig: AppConfig = {
    graphqlEndpoint: "http://localhost:5000/query",
    graphqlWSEndpoint: "ws://localhost:5000/query",
    fileEndpoint: "http://localhost:5000/files/",
    defaultResultsPerPage: 20,
  };