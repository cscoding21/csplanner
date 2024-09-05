import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "http://localhost:5000/query",
  documents: "src/lib/graphql/schema/*.graphql",
  generates: {
    "src/lib/graphql/generated/sdk.ts": {
      config: {
        useTypeImports: true,
      },
      plugins: [
        "typescript", 
        "typescript-generic-sdk",
      ]
    },
    "./graphql.schema.json": {
      plugins: ["introspection"]
    }
  }
};

export default config;