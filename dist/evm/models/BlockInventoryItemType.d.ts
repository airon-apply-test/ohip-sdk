/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { BlockInventoryItemSourceType } from './BlockInventoryItemSourceType';
import type { DateRangeType } from './DateRangeType';
import type { ItemInfoType } from './ItemInfoType';
/**
 * An identifier used to uniquely reference an object in a system (e.g. an airline reservation reference, customer profile reference, booking confirmation number, or a reference to a previous availability quote).
 * @export
 * @interface BlockInventoryItemType
 */
export interface BlockInventoryItemType {
    /**
     * URL that identifies the location associated with the record identified by the UniqueID.
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    url?: string;
    /**
     * A reference to the type of object defined by the UniqueID elementSpace.
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    type?: string;
    /**
     * The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    instance?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    idContext?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    id?: string;
    /**
     * Additional identifying value assigned by the creating system.
     * @type {number}
     * @memberof BlockInventoryItemType
     */
    idExtension?: number;
    /**
     *
     * @type {ItemInfoType}
     * @memberof BlockInventoryItemType
     */
    item?: ItemInfoType;
    /**
     * Number of items booked.
     * @type {number}
     * @memberof BlockInventoryItemType
     */
    quantity?: number;
    /**
     *
     * @type {DateRangeType}
     * @memberof BlockInventoryItemType
     */
    blockDates?: DateRangeType;
    /**
     *
     * @type {BlockInventoryItemSourceType}
     * @memberof BlockInventoryItemType
     */
    source?: BlockInventoryItemSourceType;
}
/**
 * Check if a given object implements the BlockInventoryItemType interface.
 */
export declare function instanceOfBlockInventoryItemType(value: object): boolean;
export declare function BlockInventoryItemTypeFromJSON(json: any): BlockInventoryItemType;
export declare function BlockInventoryItemTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockInventoryItemType;
export declare function BlockInventoryItemTypeToJSON(value?: BlockInventoryItemType | null): any;
