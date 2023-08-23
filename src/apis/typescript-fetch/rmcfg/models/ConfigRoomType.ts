/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { HousekeepingCreditsType } from './HousekeepingCreditsType';
import {
    HousekeepingCreditsTypeFromJSON,
    HousekeepingCreditsTypeFromJSONTyped,
    HousekeepingCreditsTypeToJSON,
} from './HousekeepingCreditsType';
import type { RatePlanRatingType } from './RatePlanRatingType';
import {
    RatePlanRatingTypeFromJSON,
    RatePlanRatingTypeFromJSONTyped,
    RatePlanRatingTypeToJSON,
} from './RatePlanRatingType';
import type { RoomComponentType } from './RoomComponentType';
import {
    RoomComponentTypeFromJSON,
    RoomComponentTypeFromJSONTyped,
    RoomComponentTypeToJSON,
} from './RoomComponentType';
import type { RoomFeatureType } from './RoomFeatureType';
import {
    RoomFeatureTypeFromJSON,
    RoomFeatureTypeFromJSONTyped,
    RoomFeatureTypeToJSON,
} from './RoomFeatureType';
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
import type { TranslationTextType2000 } from './TranslationTextType2000';
import {
    TranslationTextType2000FromJSON,
    TranslationTextType2000FromJSONTyped,
    TranslationTextType2000ToJSON,
} from './TranslationTextType2000';

/**
 * This type represents the primary room attributes.
 * @export
 * @interface ConfigRoomType
 */
export interface ConfigRoomType {
    /**
     * 
     * @type {RoomTypeShortInfoType}
     * @memberof ConfigRoomType
     */
    roomType?: RoomTypeShortInfoType;
    /**
     * Floor of the Room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    floor?: string;
    /**
     * Description for the Floor of the Room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    floorDescription?: string;
    /**
     * A recurring element that identifies the room features.
     * @type {Array<RoomFeatureType>}
     * @memberof ConfigRoomType
     */
    roomFeatures?: Array<RoomFeatureType>;
    /**
     * Detail Long Description Of The Room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    roomDescription?: string;
    /**
     * 
     * @type {TranslationTextType2000}
     * @memberof ConfigRoomType
     */
    description?: TranslationTextType2000;
    /**
     * This indicates room smoking preference.
     * @type {string}
     * @memberof ConfigRoomType
     */
    smokingPreference?: string;
    /**
     * This indicates the description of the room smoking preference.
     * @type {string}
     * @memberof ConfigRoomType
     */
    smokingPreferenceDescription?: string;
    /**
     * Building associated with the room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    building?: string;
    /**
     * 
     * @type {RatePlanRatingType}
     * @memberof ConfigRoomType
     */
    roomAssignmentRating?: RatePlanRatingType;
    /**
     * Indicates whether the room is accessibility compliant.
     * @type {boolean}
     * @memberof ConfigRoomType
     */
    accessible?: boolean;
    /**
     * Code of the room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    roomId?: string;
    /**
     * Indicates whether the room is a Meeting Room
     * @type {boolean}
     * @memberof ConfigRoomType
     */
    meetingRoom?: boolean;
    /**
     * Component of a room.
     * @type {Array<RoomComponentType>}
     * @memberof ConfigRoomType
     */
    roomComponents?: Array<RoomComponentType>;
    /**
     * Collection of rooms.
     * @type {Array<RoomRoomType>}
     * @memberof ConfigRoomType
     */
    connectingRooms?: Array<RoomRoomType>;
    /**
     * Published rate code of a room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    rateCode?: string;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof ConfigRoomType
     */
    rateAmount?: CurrencyAmountType;
    /**
     * Maximum occupancy of a room.
     * @type {number}
     * @memberof ConfigRoomType
     */
    maximumOccupancy?: number;
    /**
     * Display sequence of a room.
     * @type {number}
     * @memberof ConfigRoomType
     */
    sellSequence?: number;
    /**
     * This indicates if room is marked as a owner room
     * @type {boolean}
     * @memberof ConfigRoomType
     */
    ownerRoom?: boolean;
    /**
     * The Unit Grade Code attached to the room
     * @type {string}
     * @memberof ConfigRoomType
     */
    unitGradeCode?: string;
    /**
     * This indicates if smoking is allowed in the room.
     * @type {boolean}
     * @memberof ConfigRoomType
     */
    smoking?: boolean;
    /**
     * Key code of a room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    keyCode?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof ConfigRoomType
     */
    keyOptions?: Array<string>;
    /**
     * Square units of a room.
     * @type {number}
     * @memberof ConfigRoomType
     */
    squareUnits?: number;
    /**
     * Turndown service flag of a room.
     * @type {boolean}
     * @memberof ConfigRoomType
     */
    turndownService?: boolean;
    /**
     * Square unit measurement of a room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    squareUnitsMeasurement?: string;
    /**
     * Square unit measurement of a room.
     * @type {string}
     * @memberof ConfigRoomType
     */
    phoneNumber?: string;
    /**
     * Housekeeping credits of the room. This may include stayover, departure, pickup and turndown credits.
     * @type {Array<HousekeepingCreditsType>}
     * @memberof ConfigRoomType
     */
    housekeepingCredit?: Array<HousekeepingCreditsType>;
    /**
     * 
     * @type {RoomSectionType}
     * @memberof ConfigRoomType
     */
    roomSection?: RoomSectionType;
    /**
     * Number of beds of the room.
     * @type {number}
     * @memberof ConfigRoomType
     */
    noOfBeds?: number;
}

/**
 * Check if a given object implements the ConfigRoomType interface.
 */
export function instanceOfConfigRoomType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ConfigRoomTypeFromJSON(json: any): ConfigRoomType {
    return ConfigRoomTypeFromJSONTyped(json, false);
}

export function ConfigRoomTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigRoomType {
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
        'roomComponents': !exists(json, 'roomComponents') ? undefined : ((json['roomComponents'] as Array<any>).map(RoomComponentTypeFromJSON)),
        'connectingRooms': !exists(json, 'connectingRooms') ? undefined : ((json['connectingRooms'] as Array<any>).map(RoomRoomTypeFromJSON)),
        'rateCode': !exists(json, 'rateCode') ? undefined : json['rateCode'],
        'rateAmount': !exists(json, 'rateAmount') ? undefined : CurrencyAmountTypeFromJSON(json['rateAmount']),
        'maximumOccupancy': !exists(json, 'maximumOccupancy') ? undefined : json['maximumOccupancy'],
        'sellSequence': !exists(json, 'sellSequence') ? undefined : json['sellSequence'],
        'ownerRoom': !exists(json, 'ownerRoom') ? undefined : json['ownerRoom'],
        'unitGradeCode': !exists(json, 'unitGradeCode') ? undefined : json['unitGradeCode'],
        'smoking': !exists(json, 'smoking') ? undefined : json['smoking'],
        'keyCode': !exists(json, 'keyCode') ? undefined : json['keyCode'],
        'keyOptions': !exists(json, 'keyOptions') ? undefined : json['keyOptions'],
        'squareUnits': !exists(json, 'squareUnits') ? undefined : json['squareUnits'],
        'turndownService': !exists(json, 'turndownService') ? undefined : json['turndownService'],
        'squareUnitsMeasurement': !exists(json, 'squareUnitsMeasurement') ? undefined : json['squareUnitsMeasurement'],
        'phoneNumber': !exists(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'housekeepingCredit': !exists(json, 'housekeepingCredit') ? undefined : ((json['housekeepingCredit'] as Array<any>).map(HousekeepingCreditsTypeFromJSON)),
        'roomSection': !exists(json, 'roomSection') ? undefined : RoomSectionTypeFromJSON(json['roomSection']),
        'noOfBeds': !exists(json, 'noOfBeds') ? undefined : json['noOfBeds'],
    };
}

export function ConfigRoomTypeToJSON(value?: ConfigRoomType | null): any {
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
        'roomComponents': value.roomComponents === undefined ? undefined : ((value.roomComponents as Array<any>).map(RoomComponentTypeToJSON)),
        'connectingRooms': value.connectingRooms === undefined ? undefined : ((value.connectingRooms as Array<any>).map(RoomRoomTypeToJSON)),
        'rateCode': value.rateCode,
        'rateAmount': CurrencyAmountTypeToJSON(value.rateAmount),
        'maximumOccupancy': value.maximumOccupancy,
        'sellSequence': value.sellSequence,
        'ownerRoom': value.ownerRoom,
        'unitGradeCode': value.unitGradeCode,
        'smoking': value.smoking,
        'keyCode': value.keyCode,
        'keyOptions': value.keyOptions,
        'squareUnits': value.squareUnits,
        'turndownService': value.turndownService,
        'squareUnitsMeasurement': value.squareUnitsMeasurement,
        'phoneNumber': value.phoneNumber,
        'housekeepingCredit': value.housekeepingCredit === undefined ? undefined : ((value.housekeepingCredit as Array<any>).map(HousekeepingCreditsTypeToJSON)),
        'roomSection': RoomSectionTypeToJSON(value.roomSection),
        'noOfBeds': value.noOfBeds,
    };
}

