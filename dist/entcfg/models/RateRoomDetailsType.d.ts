/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { MasterRestrictionStatusesType } from './MasterRestrictionStatusesType';
import type { MembershipSearchType } from './MembershipSearchType';
import type { RoomStayType } from './RoomStayType';
/**
 *
 * @export
 * @interface RateRoomDetailsType
 */
export interface RateRoomDetailsType {
    /**
     * Detail regarding customer loyalty program.
     * @type {Array<MembershipSearchType>}
     * @memberof RateRoomDetailsType
     */
    memberships?: Array<MembershipSearchType>;
    /**
     *
     * @type {MasterRestrictionStatusesType}
     * @memberof RateRoomDetailsType
     */
    restrictionType?: MasterRestrictionStatusesType;
    /**
     * Room stay information.
     * @type {Array<RoomStayType>}
     * @memberof RateRoomDetailsType
     */
    roomStays?: Array<RoomStayType>;
}
/**
 * Check if a given object implements the RateRoomDetailsType interface.
 */
export declare function instanceOfRateRoomDetailsType(value: object): boolean;
export declare function RateRoomDetailsTypeFromJSON(json: any): RateRoomDetailsType;
export declare function RateRoomDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RateRoomDetailsType;
export declare function RateRoomDetailsTypeToJSON(value?: RateRoomDetailsType | null): any;
