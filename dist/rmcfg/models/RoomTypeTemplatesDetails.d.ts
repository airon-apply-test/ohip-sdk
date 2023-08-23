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
import type { InstanceLink } from './InstanceLink';
import type { RoomTypeSummaryType } from './RoomTypeSummaryType';
import type { RoomTypeType } from './RoomTypeType';
import type { WarningType } from './WarningType';
/**
 * Response object for information regarding room type template of a property.
 * @export
 * @interface RoomTypeTemplatesDetails
 */
export interface RoomTypeTemplatesDetails {
    /**
     * This type holds collection of room type.
     * @type {Array<RoomTypeSummaryType>}
     * @memberof RoomTypeTemplatesDetails
     */
    roomTypeTemplatesSummary?: Array<RoomTypeSummaryType>;
    /**
     * This type holds collection of room type.
     * @type {Array<RoomTypeType>}
     * @memberof RoomTypeTemplatesDetails
     */
    roomTypeTemplates?: Array<RoomTypeType>;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof RoomTypeTemplatesDetails
     */
    totalPages?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof RoomTypeTemplatesDetails
     */
    offset?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof RoomTypeTemplatesDetails
     */
    limit?: number;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof RoomTypeTemplatesDetails
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof RoomTypeTemplatesDetails
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof RoomTypeTemplatesDetails
     */
    count?: number;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof RoomTypeTemplatesDetails
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof RoomTypeTemplatesDetails
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the RoomTypeTemplatesDetails interface.
 */
export declare function instanceOfRoomTypeTemplatesDetails(value: object): boolean;
export declare function RoomTypeTemplatesDetailsFromJSON(json: any): RoomTypeTemplatesDetails;
export declare function RoomTypeTemplatesDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomTypeTemplatesDetails;
export declare function RoomTypeTemplatesDetailsToJSON(value?: RoomTypeTemplatesDetails | null): any;
