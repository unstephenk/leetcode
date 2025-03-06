import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class ItemService {
  private items: string[] = ['Item 1', 'Item 2', 'Item 3', 'Item 4', 'Item 5'];

  constructor() { }

  getItems(): string[] {
    return this.items;
  }

  addItems(item: string): void {
    this.items.push(item);
  }
}
