import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";


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
        return this.http.post(`${this.API}/${name}/start`, {});
    }

    stop(name: string): Observable<any> {
        return this.http.post(`${this.API}/${name}/stop`, {});
    }
}