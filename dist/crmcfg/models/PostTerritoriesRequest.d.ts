/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InstanceLink } from './InstanceLink';
import type { TerritoryType } from './TerritoryType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PostTerritoriesRequest
 */
export interface PostTerritoriesRequest {
    /**
     * List of Territories.
     * @type {Array<TerritoryType>}
     * @memberof PostTerritoriesRequest
     */
    territories?: Array<TerritoryType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PostTerritoriesRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostTerritoriesRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostTerritoriesRequest interface.
 */
export declare function instanceOfPostTerritoriesRequest(value: object): boolean;
export declare function PostTerritoriesRequestFromJSON(json: any): PostTerritoriesRequest;
export declare function PostTerritoriesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostTerritoriesRequest;
export declare function PostTerritoriesRequestToJSON(value?: PostTerritoriesRequest | null): any;
