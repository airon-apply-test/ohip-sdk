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
import type { CommentInfoType } from './CommentInfoType';
/**
 * List of Notes for the customer.
 * @export
 * @interface ProfileTypeComments
 */
export interface ProfileTypeComments {
    /**
     * Collection of Detailed information on comments for the customer.
     * @type {Array<CommentInfoType>}
     * @memberof ProfileTypeComments
     */
    commentInfo?: Array<CommentInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof ProfileTypeComments
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof ProfileTypeComments
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof ProfileTypeComments
     */
    count?: number;
}
/**
 * Check if a given object implements the ProfileTypeComments interface.
 */
export declare function instanceOfProfileTypeComments(value: object): boolean;
export declare function ProfileTypeCommentsFromJSON(json: any): ProfileTypeComments;
export declare function ProfileTypeCommentsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTypeComments;
export declare function ProfileTypeCommentsToJSON(value?: ProfileTypeComments | null): any;
