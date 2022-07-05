export interface List {
    id: number,
    imageURL?: string,
    name: string,
    ownerId: number,
    url: string
}

export interface Task  {
    id: number,
    createdById: number;
    listId: number,
    title: string,
    description: string,
    url: string,
    isDone: boolean,
    createdAt: string
}