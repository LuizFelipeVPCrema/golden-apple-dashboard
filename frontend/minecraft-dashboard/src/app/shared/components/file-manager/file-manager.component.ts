import { Component, Input, OnInit } from '@angular/core';
import { ServerService } from '../../../core/services/server.service';
import { CommonModule } from '@angular/common';


@Component({
  selector: 'app-file-manager',
  templateUrl: './file-manager.component.html',
  styleUrls: ['./file-manager.component.scss'],
  imports: [CommonModule]
})
export class FileManagerComponent implements OnInit {
  @Input() serverId!: string;
  files: string[] = [];

  constructor(private serverService: ServerService) {}

  async ngOnInit() {
    this.files = await this.serverService.getServerFiles(this.serverId);
  }
}
