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
import type { HousekeepingSectionType } from './HousekeepingSectionType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PostHousekeepingSectionsRequest
 */
export interface PostHousekeepingSectionsRequest {
    /**
     * List of the housekeeping sections to be configured or fetched
     * @type {Array<HousekeepingSectionType>}
     * @memberof PostHousekeepingSectionsRequest
     */
    housekeepingSections?: Array<HousekeepingSectionType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PostHousekeepingSectionsRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostHousekeepingSectionsRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostHousekeepingSectionsRequest interface.
 */
export declare function instanceOfPostHousekeepingSectionsRequest(value: object): boolean;
export declare function PostHousekeepingSectionsRequestFromJSON(json: any): PostHousekeepingSectionsRequest;
export declare function PostHousekeepingSectionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostHousekeepingSectionsRequest;
export declare function PostHousekeepingSectionsRequestToJSON(value?: PostHousekeepingSectionsRequest | null): any;
