import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-server-creator',
  templateUrl: './server-creator.component.html',
  styleUrls: ['./server-creator.component.scss'],
  imports: [CommonModule, FormsModule]
})
export class ServerCreatorComponent {
  server = {
    name: '',
    port: 25565,
    project: '',
    gameMode: 'survival',
    levelType: 'default'
  };

  gameModes = ['survival', 'creative', 'adventure', 'spectator'];
  levelTypes = ['default', 'flat', 'largeBiomes', 'amplified', 'buffet'];

  constructor(private router: Router) {}

  submitServer() {
    // Aqui futuramente será chamado o service que envia os dados para o backend
    console.log('Servidor criado:', this.server);
    // Redirecionar para a lista ou dashboard após criação
    this.router.navigate(['/dashboard']);
  }
}
