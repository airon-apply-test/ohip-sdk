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
import type { StayFutureListType } from './StayFutureListType';
import type { StayHistoryListType } from './StayHistoryListType';
/**
 * Information of History and Future Reservation details attached to Profiles.
 * @export
 * @interface ReservationStayHistoryFutureInfoType
 */
export interface ReservationStayHistoryFutureInfoType {
    /**
     *
     * @type {StayHistoryListType}
     * @memberof ReservationStayHistoryFutureInfoType
     */
    historyList?: StayHistoryListType;
    /**
     *
     * @type {StayFutureListType}
     * @memberof ReservationStayHistoryFutureInfoType
     */
    futureList?: StayFutureListType;
}
/**
 * Check if a given object implements the ReservationStayHistoryFutureInfoType interface.
 */
export declare function instanceOfReservationStayHistoryFutureInfoType(value: object): boolean;
export declare function ReservationStayHistoryFutureInfoTypeFromJSON(json: any): ReservationStayHistoryFutureInfoType;
export declare function ReservationStayHistoryFutureInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationStayHistoryFutureInfoType;
export declare function ReservationStayHistoryFutureInfoTypeToJSON(value?: ReservationStayHistoryFutureInfoType | null): any;
