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
import type { TravelAgentTypeType } from './TravelAgentTypeType';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PostTravelAgentTypesRequest
 */
export interface PostTravelAgentTypesRequest {
    /**
     * List of Travel Agent Types.
     * @type {Array<TravelAgentTypeType>}
     * @memberof PostTravelAgentTypesRequest
     */
    travelAgentTypes?: Array<TravelAgentTypeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PostTravelAgentTypesRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostTravelAgentTypesRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PostTravelAgentTypesRequest interface.
 */
export declare function instanceOfPostTravelAgentTypesRequest(value: object): boolean;
export declare function PostTravelAgentTypesRequestFromJSON(json: any): PostTravelAgentTypesRequest;
export declare function PostTravelAgentTypesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostTravelAgentTypesRequest;
export declare function PostTravelAgentTypesRequestToJSON(value?: PostTravelAgentTypesRequest | null): any;
