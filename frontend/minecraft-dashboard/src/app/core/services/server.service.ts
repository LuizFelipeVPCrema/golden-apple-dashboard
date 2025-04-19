import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { firstValueFrom, Observable } from "rxjs";


export interface ServerInfo {
    name: string;
    port: string;
    status: 'running' | 'stopped';
}

@Injectable({ providedIn: 'root'})
export class ServerService {
    private readonly API = '/api/servers';

    constructor(private http: HttpClient) {}

    list(): Observable<ServerInfo[]> {
        return this.http.get<ServerInfo[]>(this.API);
    }

    delete(name: string): Observable<any> {
        return this.http.delete(`${this.API}/${name}`);
    }

    start(name: string): Observable<any> {
        return this.http.put(`${this.API}/${name}/start`, {});
    }

    stop(name: string): Observable<any> {
        return this.http.put(`${this.API}/${name}/stop`, {});
    }

    /**
   * Busca os detalhes de um servidor específico pelo ID
   * @param id ID do servidor
   * @returns Promise com os dados do servidor
   */
  async getServerById(id: string): Promise<any> {
    return await firstValueFrom(this.http.get<any>(`${this.API}/${id}`));
  }

  /**
   * Lista os arquivos de um servidor pelo ID
   * @param id ID do servidor
   * @returns Promise com uma lista de arquivos
   */
  async getServerFiles(id: string): Promise<string[]> {
    return await firstValueFrom(this.http.get<string[]>(`${this.API}/${id}/files`));
  }
}