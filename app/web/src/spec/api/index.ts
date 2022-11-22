import axios, { AxiosInstance } from 'axios';

interface IApi {
  getMessage: () => Promise<string>;
}

type ApiConfig = {
  basePath?: string;
};
export class Api implements IApi {
  private readonly ax: AxiosInstance;
  constructor(config: ApiConfig) {
    this.ax = axios.create({
      baseURL: config.basePath,
    });
  }

  async getMessage() {
    const res = await this.ax.get('/');
    return res.data;
  }
}
