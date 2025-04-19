import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-console',
  templateUrl: './console.component.html',
  styleUrls: ['./console.component.scss'],
  imports: [CommonModule]
})
export class ConsoleComponent {
  @Input() serverId!: string;
  logs: string[] = [];

  sendCommand(command: string) {
    // Emitir comando via WebSocket (ex: serverService.sendCommand())
    this.logs.push(`> ${command}`);
  }

  onEnter(event: KeyboardEvent) {
    const input = event.target as HTMLInputElement;
    const value = input.value.trim();
    if (value) {
      this.sendCommand(value);
      input.value = '';
    }
  }
}
