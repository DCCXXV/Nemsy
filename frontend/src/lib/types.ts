export interface User {
	ID: number;
	StudyID: number | null;
	GoogleSub: string;
	Email: string;
	FullName: string;
	Pfp: string;
	Hd: string;
	CreatedAt: string;
}

export interface Study {
	ID: number;
	Name: string;
}
