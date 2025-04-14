import { Component, EventEmitter, Input, Output } from '@angular/core';
import { ServerInfo } from '../../../core/services/server.service';
import { CommonModule } from '@angular/common';

@Component({
    selector: 'app-server-card',
    templateUrl: './server-card.component.html',
    styleUrls: ['./server-card.component.scss'],
    imports: [CommonModule]
})
export class ServerCardComponent {
    @Input() server!: ServerInfo;
    @Output() delete = new EventEmitter<string>();
    @Output() start = new EventEmitter<string>();
    @Output() stop = new EventEmitter<string>();
}