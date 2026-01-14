export class Activity {
    activityName: string = '';
    activityResult: string = ''
    activityDate: string = '';
}

export interface ActivityResponse {
    activities: string[];
}