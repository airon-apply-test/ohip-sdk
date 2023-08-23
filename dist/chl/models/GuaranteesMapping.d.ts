/**
 * OPERA Cloud Channel Configuration API
 * APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { GuaranteeMappingType } from './GuaranteeMappingType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Response object for fetching external system guarantees.
 * @export
 * @interface GuaranteesMapping
 */
export interface GuaranteesMapping {
    /**
     * Information about an external system guarantee mapping.
     * @type {Array<GuaranteeMappingType>}
     * @memberof GuaranteesMapping
     */
    guaranteesMapping?: Array<GuaranteeMappingType>;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof GuaranteesMapping
     */
    totalPages?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof GuaranteesMapping
     */
    offset?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof GuaranteesMapping
     */
    limit?: number;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof GuaranteesMapping
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof GuaranteesMapping
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof GuaranteesMapping
     */
    count?: number;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof GuaranteesMapping
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof GuaranteesMapping
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the GuaranteesMapping interface.
 */
export declare function instanceOfGuaranteesMapping(value: object): boolean;
export declare function GuaranteesMappingFromJSON(json: any): GuaranteesMapping;
export declare function GuaranteesMappingFromJSONTyped(json: any, ignoreDiscriminator: boolean): GuaranteesMapping;
export declare function GuaranteesMappingToJSON(value?: GuaranteesMapping | null): any;
