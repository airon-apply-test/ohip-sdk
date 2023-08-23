/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CurrentServicingAttendantType } from './CurrentServicingAttendantType';
import {
    CurrentServicingAttendantTypeFromJSON,
    CurrentServicingAttendantTypeFromJSONTyped,
    CurrentServicingAttendantTypeToJSON,
} from './CurrentServicingAttendantType';
import type { HousekeepingType } from './HousekeepingType';
import {
    HousekeepingTypeFromJSON,
    HousekeepingTypeFromJSONTyped,
    HousekeepingTypeToJSON,
} from './HousekeepingType';
import type { RatePlanRatingType } from './RatePlanRatingType';
import {
    RatePlanRatingTypeFromJSON,
    RatePlanRatingTypeFromJSONTyped,
    RatePlanRatingTypeToJSON,
} from './RatePlanRatingType';
import type { ReservationShortInfoType } from './ReservationShortInfoType';
import {
    ReservationShortInfoTypeFromJSON,
    ReservationShortInfoTypeFromJSONTyped,
    ReservationShortInfoTypeToJSON,
} from './ReservationShortInfoType';
import type { RoomDiscrepancyType } from './RoomDiscrepancyType';
import {
    RoomDiscrepancyTypeFromJSON,
    RoomDiscrepancyTypeFromJSONTyped,
    RoomDiscrepancyTypeToJSON,
} from './RoomDiscrepancyType';
import type { RoomFeatureType } from './RoomFeatureType';
import {
    RoomFeatureTypeFromJSON,
    RoomFeatureTypeFromJSONTyped,
    RoomFeatureTypeToJSON,
} from './RoomFeatureType';
import type { RoomHoldType } from './RoomHoldType';
import {
    RoomHoldTypeFromJSON,
    RoomHoldTypeFromJSONTyped,
    RoomHoldTypeToJSON,
} from './RoomHoldType';
import type { RoomOutOfOrderType } from './RoomOutOfOrderType';
import {
    RoomOutOfOrderTypeFromJSON,
    RoomOutOfOrderTypeFromJSONTyped,
    RoomOutOfOrderTypeToJSON,
} from './RoomOutOfOrderType';
import type { RoomRoomType } from './RoomRoomType';
import {
    RoomRoomTypeFromJSON,
    RoomRoomTypeFromJSONTyped,
    RoomRoomTypeToJSON,
} from './RoomRoomType';
import type { RoomSectionType } from './RoomSectionType';
import {
    RoomSectionTypeFromJSON,
    RoomSectionTypeFromJSONTyped,
    RoomSectionTypeToJSON,
} from './RoomSectionType';
import type { RoomTypeShortInfoType } from './RoomTypeShortInfoType';
import {
    RoomTypeShortInfoTypeFromJSON,
    RoomTypeShortInfoTypeFromJSONTyped,
    RoomTypeShortInfoTypeToJSON,
} from './RoomTypeShortInfoType';
import type { SitePlanSectionType } from './SitePlanSectionType';
import {
    SitePlanSectionTypeFromJSON,
    SitePlanSectionTypeFromJSONTyped,
    SitePlanSectionTypeToJSON,
} from './SitePlanSectionType';
import type { TranslationTextType2000 } from './TranslationTextType2000';
import {
    TranslationTextType2000FromJSON,
    TranslationTextType2000FromJSONTyped,
    TranslationTextType2000ToJSON,
} from './TranslationTextType2000';
import type { TurndownInfoType } from './TurndownInfoType';
import {
    TurndownInfoTypeFromJSON,
    TurndownInfoTypeFromJSONTyped,
    TurndownInfoTypeToJSON,
} from './TurndownInfoType';

/**
 * This type represents the primary room attributes.
 * @export
 * @interface RoomType
 */
export interface RoomType {
    /**
     * 
     * @type {RoomTypeShortInfoType}
     * @memberof RoomType
     */
    roomType?: RoomTypeShortInfoType;
    /**
     * Floor of the Room.
     * @type {string}
     * @memberof RoomType
     */
    floor?: string;
    /**
     * Description for the Floor of the Room.
     * @type {string}
     * @memberof RoomType
     */
    floorDescription?: string;
    /**
     * A recurring element that identifies the room features.
     * @type {Array<RoomFeatureType>}
     * @memberof RoomType
     */
    roomFeatures?: Array<RoomFeatureType>;
    /**
     * Detail Long Description Of The Room.
     * @type {string}
     * @memberof RoomType
     */
    roomDescription?: string;
    /**
     * 
     * @type {TranslationTextType2000}
     * @memberof RoomType
     */
    description?: TranslationTextType2000;
    /**
     * This indicates room smoking preference.
     * @type {string}
     * @memberof RoomType
     */
    smokingPreference?: string;
    /**
     * This indicates the description of the room smoking preference.
     * @type {string}
     * @memberof RoomType
     */
    smokingPreferenceDescription?: string;
    /**
     * Building associated with the room.
     * @type {string}
     * @memberof RoomType
     */
    building?: string;
    /**
     * 
     * @type {RatePlanRatingType}
     * @memberof RoomType
     */
    roomAssignmentRating?: RatePlanRatingType;
    /**
     * Indicates whether the room is accessibility compliant.
     * @type {boolean}
     * @memberof RoomType
     */
    accessible?: boolean;
    /**
     * Code of the room.
     * @type {string}
     * @memberof RoomType
     */
    roomId?: string;
    /**
     * Indicates whether the room is a Meeting Room
     * @type {boolean}
     * @memberof RoomType
     */
    meetingRoom?: boolean;
    /**
     * 
     * @type {RoomSectionType}
     * @memberof RoomType
     */
    roomSection?: RoomSectionType;
    /**
     * 
     * @type {HousekeepingType}
     * @memberof RoomType
     */
    housekeeping?: HousekeepingType;
    /**
     * Date Range and reasons for a room being Out of Order/Out Of Service (If the room is OO/OS).
     * @type {Array<RoomOutOfOrderType>}
     * @memberof RoomType
     */
    outOfOrder?: Array<RoomOutOfOrderType>;
    /**
     * Room Discrepancies between front office and housekeeping room status and number of persons in the room.
     * @type {Array<RoomDiscrepancyType>}
     * @memberof RoomType
     */
    discrepancy?: Array<RoomDiscrepancyType>;
    /**
     * 
     * @type {SitePlanSectionType}
     * @memberof RoomType
     */
    sitePlanSection?: SitePlanSectionType;
    /**
     * This flag indicates component room.
     * @type {string}
     * @memberof RoomType
     */
    componentRoomNumber?: string;
    /**
     * Collection of rooms.
     * @type {Array<RoomRoomType>}
     * @memberof RoomType
     */
    connectingRooms?: Array<RoomRoomType>;
    /**
     * Collection of rooms.
     * @type {Array<RoomRoomType>}
     * @memberof RoomType
     */
    componentRooms?: Array<RoomRoomType>;
    /**
     * 
     * @type {Array<CurrentServicingAttendantType>}
     * @memberof RoomType
     */
    attendant?: Array<CurrentServicingAttendantType>;
    /**
     * 
     * @type {RoomHoldType}
     * @memberof RoomType
     */
    hold?: RoomHoldType;
    /**
     * 
     * @type {TurndownInfoType}
     * @memberof RoomType
     */
    turndownInfo?: TurndownInfoType;
    /**
     * 
     * @type {Array<ReservationShortInfoType>}
     * @memberof RoomType
     */
    resvInfo?: Array<ReservationShortInfoType>;
    /**
     * List of component room numbers.
     * @type {Array<string>}
     * @memberof RoomType
     */
    componentRoomNumbers?: Array<string>;
}

/**
 * Check if a given object implements the RoomType interface.
 */
export function instanceOfRoomType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoomTypeFromJSON(json: any): RoomType {
    return RoomTypeFromJSONTyped(json, false);
}

export function RoomTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomType': !exists(json, 'roomType') ? undefined : RoomTypeShortInfoTypeFromJSON(json['roomType']),
        'floor': !exists(json, 'floor') ? undefined : json['floor'],
        'floorDescription': !exists(json, 'floorDescription') ? undefined : json['floorDescription'],
        'roomFeatures': !exists(json, 'roomFeatures') ? undefined : ((json['roomFeatures'] as Array<any>).map(RoomFeatureTypeFromJSON)),
        'roomDescription': !exists(json, 'roomDescription') ? undefined : json['roomDescription'],
        'description': !exists(json, 'description') ? undefined : TranslationTextType2000FromJSON(json['description']),
        'smokingPreference': !exists(json, 'smokingPreference') ? undefined : json['smokingPreference'],
        'smokingPreferenceDescription': !exists(json, 'smokingPreferenceDescription') ? undefined : json['smokingPreferenceDescription'],
        'building': !exists(json, 'building') ? undefined : json['building'],
        'roomAssignmentRating': !exists(json, 'roomAssignmentRating') ? undefined : RatePlanRatingTypeFromJSON(json['roomAssignmentRating']),
        'accessible': !exists(json, 'accessible') ? undefined : json['accessible'],
        'roomId': !exists(json, 'roomId') ? undefined : json['roomId'],
        'meetingRoom': !exists(json, 'meetingRoom') ? undefined : json['meetingRoom'],
        'roomSection': !exists(json, 'roomSection') ? undefined : RoomSectionTypeFromJSON(json['roomSection']),
        'housekeeping': !exists(json, 'housekeeping') ? undefined : HousekeepingTypeFromJSON(json['housekeeping']),
        'outOfOrder': !exists(json, 'outOfOrder') ? undefined : ((json['outOfOrder'] as Array<any>).map(RoomOutOfOrderTypeFromJSON)),
        'discrepancy': !exists(json, 'discrepancy') ? undefined : ((json['discrepancy'] as Array<any>).map(RoomDiscrepancyTypeFromJSON)),
        'sitePlanSection': !exists(json, 'sitePlanSection') ? undefined : SitePlanSectionTypeFromJSON(json['sitePlanSection']),
        'componentRoomNumber': !exists(json, 'componentRoomNumber') ? undefined : json['componentRoomNumber'],
        'connectingRooms': !exists(json, 'connectingRooms') ? undefined : ((json['connectingRooms'] as Array<any>).map(RoomRoomTypeFromJSON)),
        'componentRooms': !exists(json, 'componentRooms') ? undefined : ((json['componentRooms'] as Array<any>).map(RoomRoomTypeFromJSON)),
        'attendant': !exists(json, 'attendant') ? undefined : ((json['attendant'] as Array<any>).map(CurrentServicingAttendantTypeFromJSON)),
        'hold': !exists(json, 'hold') ? undefined : RoomHoldTypeFromJSON(json['hold']),
        'turndownInfo': !exists(json, 'turndownInfo') ? undefined : TurndownInfoTypeFromJSON(json['turndownInfo']),
        'resvInfo': !exists(json, 'resvInfo') ? undefined : ((json['resvInfo'] as Array<any>).map(ReservationShortInfoTypeFromJSON)),
        'componentRoomNumbers': !exists(json, 'componentRoomNumbers') ? undefined : json['componentRoomNumbers'],
    };
}

export function RoomTypeToJSON(value?: RoomType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomType': RoomTypeShortInfoTypeToJSON(value.roomType),
        'floor': value.floor,
        'floorDescription': value.floorDescription,
        'roomFeatures': value.roomFeatures === undefined ? undefined : ((value.roomFeatures as Array<any>).map(RoomFeatureTypeToJSON)),
        'roomDescription': value.roomDescription,
        'description': TranslationTextType2000ToJSON(value.description),
        'smokingPreference': value.smokingPreference,
        'smokingPreferenceDescription': value.smokingPreferenceDescription,
        'building': value.building,
        'roomAssignmentRating': RatePlanRatingTypeToJSON(value.roomAssignmentRating),
        'accessible': value.accessible,
        'roomId': value.roomId,
        'meetingRoom': value.meetingRoom,
        'roomSection': RoomSectionTypeToJSON(value.roomSection),
        'housekeeping': HousekeepingTypeToJSON(value.housekeeping),
        'outOfOrder': value.outOfOrder === undefined ? undefined : ((value.outOfOrder as Array<any>).map(RoomOutOfOrderTypeToJSON)),
        'discrepancy': value.discrepancy === undefined ? undefined : ((value.discrepancy as Array<any>).map(RoomDiscrepancyTypeToJSON)),
        'sitePlanSection': SitePlanSectionTypeToJSON(value.sitePlanSection),
        'componentRoomNumber': value.componentRoomNumber,
        'connectingRooms': value.connectingRooms === undefined ? undefined : ((value.connectingRooms as Array<any>).map(RoomRoomTypeToJSON)),
        'componentRooms': value.componentRooms === undefined ? undefined : ((value.componentRooms as Array<any>).map(RoomRoomTypeToJSON)),
        'attendant': value.attendant === undefined ? undefined : ((value.attendant as Array<any>).map(CurrentServicingAttendantTypeToJSON)),
        'hold': RoomHoldTypeToJSON(value.hold),
        'turndownInfo': TurndownInfoTypeToJSON(value.turndownInfo),
        'resvInfo': value.resvInfo === undefined ? undefined : ((value.resvInfo as Array<any>).map(ReservationShortInfoTypeToJSON)),
        'componentRoomNumbers': value.componentRoomNumbers,
    };
}

