/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ApplicationUserType } from './ApplicationUserType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import type { TrackItActionType } from './TrackItActionType';
import type { TrackItGroup } from './TrackItGroup';
import type { TrackItId } from './TrackItId';
import type { TrackItLogType } from './TrackItLogType';
import type { TrackItReservationInfoType } from './TrackItReservationInfoType';
import type { TrackItType } from './TrackItType';
/**
 * Detailed information of a Track It item.
 * @export
 * @interface TrackItItemType
 */
export interface TrackItItemType {
    /**
     *
     * @type {string}
     * @memberof TrackItItemType
     */
    hotelId?: string;
    /**
     *
     * @type {TrackItId}
     * @memberof TrackItItemType
     */
    trackItId?: TrackItId;
    /**
     *
     * @type {TrackItGroup}
     * @memberof TrackItItemType
     */
    group?: TrackItGroup;
    /**
     *
     * @type {string}
     * @memberof TrackItItemType
     */
    ticketNumber?: string;
    /**
     *
     * @type {string}
     * @memberof TrackItItemType
     */
    referenceNumber?: string;
    /**
     *
     * @type {TrackItType}
     * @memberof TrackItItemType
     */
    type?: TrackItType;
    /**
     *
     * @type {TrackItActionType}
     * @memberof TrackItItemType
     */
    action?: TrackItActionType;
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof TrackItItemType
     */
    location?: CodeDescriptionType;
    /**
     *
     * @type {number}
     * @memberof TrackItItemType
     */
    quantity?: number;
    /**
     *
     * @type {Date}
     * @memberof TrackItItemType
     */
    followUpDate?: Date;
    /**
     *
     * @type {string}
     * @memberof TrackItItemType
     */
    description?: string;
    /**
     *
     * @type {ApplicationUserType}
     * @memberof TrackItItemType
     */
    assignedTo?: ApplicationUserType;
    /**
     *
     * @type {TrackItReservationInfoType}
     * @memberof TrackItItemType
     */
    reservationInfo?: TrackItReservationInfoType;
    /**
     *
     * @type {Array<TrackItLogType>}
     * @memberof TrackItItemType
     */
    trackItLogList?: Array<TrackItLogType>;
}
/**
 * Check if a given object implements the TrackItItemType interface.
 */
export declare function instanceOfTrackItItemType(value: object): boolean;
export declare function TrackItItemTypeFromJSON(json: any): TrackItItemType;
export declare function TrackItItemTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TrackItItemType;
export declare function TrackItItemTypeToJSON(value?: TrackItItemType | null): any;
