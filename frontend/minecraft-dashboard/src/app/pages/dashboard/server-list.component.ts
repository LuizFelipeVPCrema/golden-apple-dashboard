import { Component, OnInit } from "@angular/core";
import { ServerInfo, ServerService } from "../../core/services/server.service";
import { ServerCardComponent } from "../../shared/components/server-card/server-card.component";
import { CommonModule } from "@angular/common";

@Component({
    selector: 'app-server-list',
    templateUrl: './server-list.component.html',
    styleUrls: ['./server-list.component.scss'],
    imports: [CommonModule, ServerCardComponent]
})
export class ServerListComponent implements OnInit {
    servers: ServerInfo[] = [];
    loading = false;

    constructor(private serverService: ServerService) {}

    ngOnInit(): void {
        this.fetchServers();
    }

    fetchServers(): void {
        this.loading = true;
        this.serverService.list().subscribe({
            next: (data) => (this.servers = data),
            complete: () => (this.loading = false),
        });
    }

    deleteServer(name: string): void {
        this.serverService.delete(name).subscribe(() => this.fetchServers());
    }

    startServer(name: string): void {
        this.serverService.start(name).subscribe(() => this.fetchServers());
    }

    stopServer(name: string): void {
        this.serverService.stop(name).subscribe(() => this.fetchServers());
    }
}