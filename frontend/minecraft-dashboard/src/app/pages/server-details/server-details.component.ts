import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ServerService } from '../../core/services/server.service';
import { CommonModule } from '@angular/common';
import { ServerInfoComponent } from "../../shared/components/server-info/server-info.component";
import { ServerConfigComponent } from "../../shared/components/server-config/server-config.component";
import { FileManagerComponent } from "../../shared/components/file-manager/file-manager.component";
import { ConsoleComponent } from "../../shared/components/console/console.component";

@Component({
  selector: 'app-server-details',
  templateUrl: './server-details.component.html',
  styleUrls: ['./server-details.component.scss'],
  imports: [CommonModule, ServerInfoComponent, ServerConfigComponent, FileManagerComponent, ConsoleComponent]
})
export class ServerDetailsComponent implements OnInit {
  serverId: string = '';
  serverInfo: any;
  loading = true;

  constructor(private route: ActivatedRoute, private serverService: ServerService) {}

  ngOnInit() {
    this.serverId = this.route.snapshot.paramMap.get('id') || '';
    this.loadServerDetails();
  }

  async loadServerDetails() {
    this.loading = true;
    this.serverInfo = await this.serverService.getServerById(this.serverId);
    this.loading = false;
  }
}
