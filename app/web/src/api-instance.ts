import { Api } from './spec/api';

const isBrowser = typeof window !== 'undefined';
const API_BASEPATH = isBrowser ? process.env.NEXT_PUBLIC_CLIENT_SIDE_API_BASEPATH : process.env.SERVER_SIDE_API_BASEPATH;

export const api = new Api({
  basePath: API_BASEPATH,
});
