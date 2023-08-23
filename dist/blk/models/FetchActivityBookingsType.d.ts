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
import type { ActivityListInner } from './ActivityListInner';
import type { AddressType } from './AddressType';
import type { PersonNameType } from './PersonNameType';
import type { UniqueIDType } from './UniqueIDType';
/**
 * Used in the request message to describe the "filtering Criteria" when executing an activity lookup.
 * @export
 * @interface FetchActivityBookingsType
 */
export interface FetchActivityBookingsType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof FetchActivityBookingsType
     */
    profileId?: Array<UniqueIDType>;
    /**
     *
     * @type {PersonNameType}
     * @memberof FetchActivityBookingsType
     */
    personName?: PersonNameType;
    /**
     *
     * @type {AddressType}
     * @memberof FetchActivityBookingsType
     */
    address?: AddressType;
    /**
     * A collection of Activity objects.
     * @type {Array<ActivityListInner>}
     * @memberof FetchActivityBookingsType
     */
    activities?: Array<ActivityListInner>;
    /**
     * Hotel Code, It is used to filter hotel specific children to this specific hotel code.
     * @type {string}
     * @memberof FetchActivityBookingsType
     */
    hotelId?: string;
}
/**
 * Check if a given object implements the FetchActivityBookingsType interface.
 */
export declare function instanceOfFetchActivityBookingsType(value: object): boolean;
export declare function FetchActivityBookingsTypeFromJSON(json: any): FetchActivityBookingsType;
export declare function FetchActivityBookingsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FetchActivityBookingsType;
export declare function FetchActivityBookingsTypeToJSON(value?: FetchActivityBookingsType | null): any;
